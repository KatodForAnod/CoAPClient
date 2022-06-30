package server

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/controller"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	controller controller.ServerController
}

func (s *Server) StartServer(config config.Config, controller controller.ServerController) {
	s.controller = controller

	http.HandleFunc("/device/metrics", s.getInformationFromIotDevice)
	http.HandleFunc("/device/add", s.addIotDevice)
	http.HandleFunc("/device/rm", s.removeIotDevice)
	http.HandleFunc("/logs", s.getLogs)

	fmt.Println("Server is listening... ", config.ProxyServerAddr)
	log.Fatal(http.ListenAndServe(config.ProxyServerAddr, nil))
}
