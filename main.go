package main

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/controller"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/logsetting"
	"CoAPProxyServer/pkg/memory"
	serv "CoAPProxyServer/pkg/server"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)
	er := logsetting.LogInit()
	if er != nil {
		log.Fatalln(er)
	}

	mem := memory.MemBuff{}
	mem.InitStruct()
	//mem := memory.MemoryFmt{}

	conf, _ := config.LoadConfig()
	iotDev := iot.IoTDevice{}
	iotDev.Init(conf.IoTsDevices[0])

	iotController := controller.IoTsController{}
	iotController.Init(conf, &mem)

	var arr []*iot.IoTDevice
	arr = append(arr, &iotDev)
	iotController.AddIoTs(arr)

	iotController.StartInformationCollect()

	controll := controller.Controller{}
	server := serv.Server{}

	controll.InitStruct(conf, &mem, iotController)
	server.StartServer(conf, controll)
	return

	//time.Sleep(time.Second * 8)
	//mem.FlushToFile(conf.IoTsDevices[0].Name)
	iotController.StopInformationCollect()
	time.Sleep(time.Second * 4)
	iotController.StartInformationCollect()
	time.Sleep(time.Second * 4)
}
