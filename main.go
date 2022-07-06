package main

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/controller"
	"CoAPProxyServer/pkg/iot"
	"CoAPProxyServer/pkg/logsetting"
	"CoAPProxyServer/pkg/memory"
	serv "CoAPProxyServer/pkg/server"
	"flag"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	var proxyServerAddr string
	flag.StringVar(&proxyServerAddr, "proxyAddr",
		"", "address of http server")
	flag.Parse()

	er := logsetting.LogInit()
	if er != nil {
		log.Fatalln(er)
	}

	conf, _ := config.LoadConfig()
	if proxyServerAddr != "" {
		conf.ProxyServerAddr = proxyServerAddr
	}

	mem := memory.MemBuff{}
	mem.InitStruct()
	//mem := memory.MemoryFmt{}

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
	server.StartServer(conf, &controll)
	return

	//time.Sleep(time.Second * 8)
	//mem.FlushToFile(conf.IoTsDevices[0].Name)
	iotController.StopInformationCollect()
	time.Sleep(time.Second * 4)
	iotController.StartInformationCollect()
	time.Sleep(time.Second * 4)
}
