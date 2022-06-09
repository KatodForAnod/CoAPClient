package controller

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/memory"
	"log"
)

type Controller struct {
	mem            memory.Memory
	ioTsController IoTsController
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

func (c *Controller) NewIotDeviceObserve(iotConfig config.IotConfig) error {
	log.Println("controller new iotDevicesObserve")
	iotDev := iot.IoTDevice{}
	iotDev.Init(iotConfig)
	var arr []*iot.IoTDevice
	arr = append(arr, &iotDev)

	err := c.ioTsController.AddIoTs(arr)
	if err != nil {
		log.Println(err)
		return err
	}

	err = c.ioTsController.StartInformationCollect()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Controller) RemoveIoTDeviceObserve(ioTsConfig []config.IotConfig) error {
	log.Println("controller remove ioTDeviceObserve")
	c.ioTsController.RemoveIoTs(ioTsConfig)
	return nil
}
