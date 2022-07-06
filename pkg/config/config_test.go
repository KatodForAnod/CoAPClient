package config

import (
	"os"
	"testing"
)

const confBody = `{
    "proxy_server_addr":"127.0.0.1:5300",
    "iots_devices":[
        {
			"addr":"127.0.0.1:5301",
			"name":"testName1"
		}, 
		{
			"addr":"127.0.0.1:5302",
			"name":"testName2"
		}
    ]
}`

func createConfig(t *testing.T) error {
	file, err := os.CreateTemp("", configPath)
	if err != nil {
		t.Error("cant create temp conf file")
		return err
	}

	configPath = file.Name()

	_, err = file.WriteString(confBody)
	if err != nil {
		t.Error("cant write to temp conf file")
		return err
	}

	return nil
}

func deleteConfig(t *testing.T) error {
	err := os.Remove(configPath)
	if err != nil {
		t.Error("cant delete temp conf file")
		return err
	}
	return nil
}

func TestLoadConfigSuccess(t *testing.T) {
	err := createConfig(t)
	if err != nil {
		t.Error(err)
		return
	}

	config, err := LoadConfig()
	if err != nil {
		t.Error(err)
		return
	}

	if config.ProxyServerAddr != "127.0.0.1:5300" {
		t.Error("wrong load file field")
		return
	}

	if len(config.IoTsDevices) != 2 {
		t.Error("wrong array size")
		return
	}

	if config.IoTsDevices[0].Name != "testName1" &&
		config.IoTsDevices[0].Addr != "127.0.0.1:5301" {
		t.Error("wrong load file field")
		return
	}

	if config.IoTsDevices[1].Name != "testName2" &&
		config.IoTsDevices[1].Addr != "127.0.0.1:5302" {
		t.Error("wrong load file field")
		return
	}

	err = deleteConfig(t)
	if err != nil {
		t.Error("cant delete temp conf file")
		return
	}
}

func TestLoadConfigFail(t *testing.T) {
	configPath = "notExist.txt"
	_, err := LoadConfig()
	if err == nil {
		t.Error("func must return error")
		return
	}
}
