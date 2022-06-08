package controller

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/iot"
	"CoAPClient/pkg/memory"
	"context"
	"log"
	"time"
)

type IoTsController struct {
	ioTDevices []iot.IoTDevice // TODO wrap into interface
	mem        memory.Memory
}

func (c *IoTsController) Init(config config.Config, mem memory.Memory) {
	c.mem = mem
}

func (c *IoTsController) AddIoTs(iots []iot.IoTDevice) {
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
		// mb create new context?
		// if device already collect inform?
		err := device.ObserveInform(ctx, c.mem.Save)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil // error only once returns
}

func (c *IoTsController) StopInformationCollect() error {
	log.Println("stop information collect")

	for _, device := range c.ioTDevices {
		err := device.StopObserveInform()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
