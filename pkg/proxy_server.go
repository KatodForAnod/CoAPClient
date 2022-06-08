package main

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/iot"
)

type Server struct {
	ioTDevices []iot.IoTDevice // TODO wrap into interface
	//connection to Another server
}

func (s *Server) StartServer(config config.Config) {

}
