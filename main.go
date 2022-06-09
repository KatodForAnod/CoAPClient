package main

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/controller"
	"CoAPClient/pkg/iot"
	"CoAPClient/pkg/memory"
	"time"
)

func main() {
	//mem := memory.MemBuff{}
	//mem.InitStruct("test.txt")
	mem := memory.MemoryFmt{}

	conf, _ := config.LoadConfig()
	iotDev := iot.IoTDevice{}
	iotDev.Init(conf.IoTsDevices[0])

	iotController := controller.IoTsController{}
	iotController.Init(conf, &mem)

	var arr []*iot.IoTDevice
	arr = append(arr, &iotDev)
	iotController.AddIoTs(arr)

	iotController.StartInformationCollect()
	time.Sleep(time.Second * 8)
	//mem.FlushToFile()

	iotController.StopInformationCollect()
	time.Sleep(time.Second * 4)
	iotController.StartInformationCollect()
	time.Sleep(time.Second * 4)
}
