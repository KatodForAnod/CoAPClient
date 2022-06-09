package controller

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/memory"
	"context"
	"github.com/plgd-dev/go-coap/v2/message"
	"log"
	"time"
)

type IoTsController struct {
	ioTDevices []*iot.IoTDevice // TODO wrap into interface
	mem        memory.Memory
}

func (c *IoTsController) Init(config config.Config, mem memory.Memory) {
	c.mem = mem
}

func (c *IoTsController) AddIoTs(iots []*iot.IoTDevice) {
	c.ioTDevices = append(c.ioTDevices, iots...)
}

func (c *IoTsController) StartInformationCollect() error {
	log.Println("start information collect")

	for _, device := range c.ioTDevices {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := device.Ping(ctx); err != nil {
			log.Println(err)
			if err := device.Connect(); err != nil {
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
		// check if already stop
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
		if err := c.mem.Save(msg, msgType); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
