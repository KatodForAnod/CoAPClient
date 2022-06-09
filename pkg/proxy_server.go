package main

import (
	"CoAPClient/pkg/config"
	"CoAPClient/pkg/controller"
	"log"
	"net/http"
)

type Server struct {
	controller controller.Controller
}

func (s *Server) getInformationFromIotDevice(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	inf, err := s.controller.GetInformation()
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(inf)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) StartServer(config config.Config, controller controller.Controller) {
	s.controller = controller

	http.HandleFunc("/", s.getInformationFromIotDevice)
	log.Fatal(http.ListenAndServe(config.ProxyServerAddr, nil))
}
