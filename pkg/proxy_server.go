package main

import (
	"CoAPProxyServer/pkg/config"
	"CoAPProxyServer/pkg/controller"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	controller controller.Controller
}

func (s *Server) getInformationFromIotDevice(w http.ResponseWriter, r *http.Request) {
	deviceNames := r.URL.Query()["deviceName"]
	if len(deviceNames) == 0 {
		log.Println("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}
	deviceName := deviceNames[0]

	inf, err := s.controller.GetInformation(deviceName)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(inf)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) StartServer(config config.Config, controller controller.Controller) {
	s.controller = controller

	http.HandleFunc("/device/metrics", s.getInformationFromIotDevice)
	log.Fatal(http.ListenAndServe(config.ProxyServerAddr, nil))
}
