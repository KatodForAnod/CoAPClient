package server

import (
	"CoAPProxyServer/pkg/config"
	"errors"
	"testing"
)

var (
	proxyServer Server
)

type Controller struct {
	ioTs []config.IotConfig
}

func (c *Controller) GetLastNRowsLogs(nRows int) ([]string, error) {
	if nRows < 0 {
		return []string{}, errors.New("wrong count rows")
	}
	return []string{"1 row", "2 row"}, nil
}

func (c *Controller) RemoveIoTDeviceObserve(ioTsConfig []config.IotConfig) error {
	return nil
}

func (c *Controller) NewIotDeviceObserve(iotConfig config.IotConfig) error {
	c.ioTs = append(c.ioTs, iotConfig)
	return nil
}

func (c *Controller) GetInformation(deviceName string) ([]byte, error) {
	for i, t := range c.ioTs {
		if t.Name == deviceName {
			return []byte{byte(i)}, nil
		}
	}

	return []byte{}, errors.New("not found")
}

func Init() {
	controller := Controller{}
	proxyServer.controller = &controller
}

func TestServer_getInformationFromIotDevice(t *testing.T) {

}

func TestServer_addIotDevice(t *testing.T) {

}

func TestServer_removeIotDevice(t *testing.T) {

}

func TestServer_getLogs(t *testing.T) {

}
