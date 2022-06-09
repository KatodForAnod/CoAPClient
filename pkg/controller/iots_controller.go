package controller

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/memory"
	"context"
	"errors"
	"github.com/plgd-dev/go-coap/v2/message"
	"log"
	"time"
)

type IoTsController struct {
	ioTDevices map[string]*iot.IoTDevice // TODO wrap into interface
	mem        memory.Memory
}

func (c *IoTsController) Init(config config.Config, mem memory.Memory) {
	c.mem = mem
	c.ioTDevices = make(map[string]*iot.IoTDevice)
}

func (c *IoTsController) AddIoTs(iots []*iot.IoTDevice) error {
	for _, device := range iots {
		if _, isExist := c.ioTDevices[device.GetName()]; isExist {
			err := errors.New("device " + device.GetName() + " already exist")
			log.Println(err)
			return err
		}
	}

	for _, device := range iots {
		c.ioTDevices[device.GetName()] = device
	}
	return nil
}

func (c *IoTsController) RemoveIoTs(IoTsConfig []config.IotConfig) {
	for _, device := range IoTsConfig {
		if iotDevice, isExist := c.ioTDevices[device.Name]; isExist {
			err := iotDevice.StopObserveInform()
			if err != nil {
				log.Println(err)
			}
			err = iotDevice.Disconnect()
			if err != nil {
				log.Println(err)
			}

			delete(c.ioTDevices, device.Name)
		}
	}
}

func (c *IoTsController) StartInformationCollect() error {
	log.Println("start information collect")

	for _, device := range c.ioTDevices {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := device.Ping(ctx); err != nil {
			log.Println(err)
			if err := device.Connect(); err != nil { // when connect need restart IsObserveInformProcess
				log.Println(err)
				continue
			}
		}

		if device.IsObserveInformProcess() {
			continue
		}

		err := device.ObserveInform(c.createSaveFunc(time.Second*2, device))
		if err != nil {
			log.Println(err)
		}
	}

	return nil // error only once returns
}

func (c *IoTsController) StopInformationCollect() error {
	log.Println("stop information collect")

	for _, device := range c.ioTDevices {
		if !device.IsObserveInformProcess() {
			continue // if device already stopped
		}
		err := device.StopObserveInform()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (c *IoTsController) createSaveFunc(d time.Duration,
	iotDevice *iot.IoTDevice) func([]byte, message.MediaType) error {
	timer := time.AfterFunc(d, func() {
		if iotDevice.IsObserveInformProcess() {
			log.Println("iot device -", iotDevice.GetName(), "not responding")
		}
	})

	return func(msg []byte, msgType message.MediaType) error {
		timer.Reset(d)
		if err := c.mem.Save(msg, msgType, iotDevice.GetName()); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
