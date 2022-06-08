package config

type Config struct {
	IoTsDevices []IotConfig `json:"io_ts_devices"`
}

type IotConfig struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

func LoadConfig() (Config, error) {
	conf := Config{IoTsDevices: []IotConfig{{
		Addr: "localhost:5688",
		Name: "testDevice",
	}}}

	return conf, nil
}
