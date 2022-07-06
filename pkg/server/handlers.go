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
	defer r.Body.Close()
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
	defer r.Body.Close()
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

func (s *Server) getLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("handler getLogs")
	countLogsArr := r.URL.Query()["countLogs"]
	if len(countLogsArr) == 0 {
		log.Println("count logs not found")
		fmt.Fprintf(w, "set count logs")
		return
	}
	countLogsStr := countLogsArr[0]
	countLogs, err := strconv.Atoi(countLogsStr)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs, err := s.controller.GetLastNRowsLogs(countLogs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allLogs := strings.Join(logs, "\n")
	_, err = w.Write([]byte(allLogs))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
