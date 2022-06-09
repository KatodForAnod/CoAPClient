package controller

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/memory"
)

type Controller struct {
	mem memory.Memory
}

func (c *Controller) InitStruct(config config.Config, mem memory.Memory) {
	c.mem = mem
}

// args to know what type info must recieve?
func (c *Controller) GetInformation() ([]byte, error) {
	return c.mem.Load()
}
