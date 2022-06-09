package config

type Config struct {
	ProxyServerAddr string      `json:"proxy_server_addr"`
	IoTsDevices     []IotConfig `json:"io_ts_devices"`
}

type IotConfig struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

func LoadConfig() (Config, error) {
	conf := Config{IoTsDevices: []IotConfig{{
		Addr: "localhost:5688",
		Name: "testDevice",
	}},
		ProxyServerAddr: "localhost:8000"}

	return conf, nil
}
