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
	log.Println("handler getInformationFromIotDevice")
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

func (s *Server) addIotDevice(w http.ResponseWriter, r *http.Request) {
	log.Println("handler addIotDevice")
	deviceNames := r.URL.Query()["deviceName"]
	if len(deviceNames) == 0 {
		log.Println("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}

	deviceAddrs := r.URL.Query()["deviceAddr"]
	if len(deviceAddrs) == 0 {
		log.Println("device addr not found")
		fmt.Fprintf(w, "set device addr")
		return
	}
	deviceName := deviceNames[0]
	deviceAddr := deviceAddrs[0]

	err := s.controller.NewIotDeviceObserve(config.IotConfig{
		Addr: deviceAddr,
		Name: deviceName,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) removeIotDevice(w http.ResponseWriter, r *http.Request) {
	log.Println("handler removeIotDevice")
	deviceNames := r.URL.Query()["deviceName"]
	if len(deviceNames) == 0 {
		log.Println("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}
	deviceName := deviceNames[0]

	err := s.controller.RemoveIoTDeviceObserve([]config.IotConfig{{Name: deviceName}})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) StartServer(config config.Config, controller controller.Controller) {
	s.controller = controller

	http.HandleFunc("/device/metrics", s.getInformationFromIotDevice)
	log.Fatal(http.ListenAndServe(config.ProxyServerAddr, nil))
}
