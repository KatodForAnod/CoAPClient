package controller

import "CoAPProxyServer/pkg/config"

type ServerController interface {
	GetLastNRowsLogs(nRows int) ([]string, error)
	RemoveIoTDeviceObserve(ioTsConfig []config.IotConfig) error
	NewIotDeviceObserve(iotConfig config.IotConfig) error
	GetInformation(deviceName string) ([]byte, error)
}
