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

// args to know what type info must recieve?
func (c *Controller) GetInformation() ([]byte, error) {
	log.Println("controller get information")
	return []byte{}, nil
}
