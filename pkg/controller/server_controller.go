package controller

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/logsetting"
	"CoAPProxyServer/pkg/memory"
)

type Controller struct {
	mem            memory.Memory
	ioTsController IoTsController
}

func (c *Controller) InitStruct(config config.Config,
	mem memory.Memory, ioTsController IoTsController) {
	c.ioTsController = ioTsController
	c.mem = mem
}

func (c *Controller) GetInformation(deviceName string) ([]byte, error) {
	load, err := c.mem.Load(deviceName)
	if err != nil {
		return []byte{}, err
	}

	return load, nil
}

func (c *Controller) NewIotDeviceObserve(iotConfig config.IotConfig) error {
	iotDev := iot.IoTDevice{}
	iotDev.Init(iotConfig)
	var arr []*iot.IoTDevice
	arr = append(arr, &iotDev)

	err := c.ioTsController.AddIoTs(arr)
	if err != nil {
		return err
	}

	err = c.ioTsController.StartInformationCollect()
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) RemoveIoTDeviceObserve(ioTsConfig []config.IotConfig) error {
	c.ioTsController.RemoveIoTs(ioTsConfig)
	return nil
}

func (c *Controller) GetLastNRowsLogs(nRows int) ([]string, error) {
	file, err := logsetting.OpenLastLogFile()
	if err != nil {
		return []string{}, err
	}

	logs, err := logsetting.GetNLastLines(file, nRows)
	if err != nil {
		return []string{}, err
	}

	return logs, nil
}
