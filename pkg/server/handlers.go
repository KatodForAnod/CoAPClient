package server

import (
	"CoAPProxyServer/pkg/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) getInformationFromIotDevice(w http.ResponseWriter, r *http.Request) {
	log.Println("handler getInformationFromIotDevice")
	defer r.Body.Close()
	deviceName := r.URL.Query().Get("deviceName")
	if deviceName == "" {
		log.Errorln("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}

	inf, err := s.controller.GetInformation(deviceName)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(inf)
	if err != nil {
		log.Errorln(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) addIotDevice(w http.ResponseWriter, r *http.Request) {
	log.Println("handler addIotDevice")
	defer r.Body.Close()
	deviceName := r.URL.Query().Get("deviceName")
	if deviceName == "" {
		log.Errorln("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}

	deviceAddr := r.URL.Query().Get("deviceAddr")
	if deviceAddr == "" {
		log.Errorln("device addr not found")
		fmt.Fprintf(w, "set device addr")
		return
	}

	err := s.controller.NewIotDeviceObserve(config.IotConfig{
		Addr: deviceAddr,
		Name: deviceName,
	})
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) removeIotDevice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("handler removeIotDevice")
	deviceName := r.URL.Query().Get("deviceName")
	if deviceName == "" {
		log.Errorln("device name not found")
		fmt.Fprintf(w, "set device name")
		return
	}

	err := s.controller.RemoveIoTDeviceObserve([]config.IotConfig{{Name: deviceName}})
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("handler getLogs")
	countLogs := r.URL.Query().Get("countLogs")
	if countLogs == "" {
		log.Errorln("count logs not found")
		fmt.Fprintf(w, "set count logs")
		return
	}
	countLogsInt, err := strconv.Atoi(countLogs)
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs, err := s.controller.GetLastNRowsLogs(countLogsInt)
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allLogs := strings.Join(logs, "\n")
	_, err = w.Write([]byte(allLogs))
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
