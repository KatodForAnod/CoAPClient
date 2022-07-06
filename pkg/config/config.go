package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	ProxyServerAddr string      `json:"proxy_server_addr"`
	IoTsDevices     []IotConfig `json:"iots_devices"`
}

type IotConfig struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

var configPath = "conf.config"

func LoadConfig() (loadedConf Config, err error) {
	/*conf := Config{IoTsDevices: []IotConfig{{
		Addr: "localhost:5688",
		Name: "testDevice",
	}},
		ProxyServerAddr: "localhost:8000"}*/
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

	err = json.Unmarshal(data, &loadedConf)
	return
}
