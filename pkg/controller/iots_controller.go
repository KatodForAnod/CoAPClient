package controller

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/iot"
	"CoAPClient/pkg/memory"
	"context"
	"github.com/plgd-dev/go-coap/v2/udp/message/pool"
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
		err := device.ObserveInform(ctx, c.saveData)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil // error only once returns
}

// if data process diff from iot device to iot device mb move func to super iot class??
func (c *IoTsController) saveData(req *pool.Message) {
	log.Printf("Got %+v\n", req)
	buff := make([]byte, 300)
	_, err := req.Body().Read(buff)
	if err != nil {
		log.Println(err)
		return
	}
	infType, err := req.Message.ContentFormat()
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(string(buff))
	if err := c.mem.Save(buff, infType); err != nil {
		log.Println(err)
		return
	}
}
