package controller

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/memory"
	"log"
)

type Controller struct {
	mem memory.Memory
}

func (c *Controller) InitStruct(config config.Config, mem memory.Memory) {
	c.mem = mem
}

func (c *Controller) GetInformation(deviceName string) ([]byte, error) {
	log.Println("controller get information of iot device", deviceName)

	load, err := c.mem.Load(deviceName)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return load, nil
}
