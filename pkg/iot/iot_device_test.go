package iot

import (
	"CoAPProxyServer/pkg/config"
	"context"
	"github.com/plgd-dev/go-coap/v2/message"
	"testing"
	"time"
)

var iotDev IoTDevice

func createDockerContainerWithIotDevice(t *testing.T) {

}

func TestIoTDevice_Init(t *testing.T) {
	iotDev = IoTDevice{}

	//createDockerContainerWithIotDevice

	conf := config.IotConfig{
		Addr: "123:8080",
		Name: "testDevice",
	}
	iotDev.Init(conf)

	if iotDev.isObserveInformProcess == nil {
		t.Error("field *bool not initialize")
		return
	}
}

func TestIoTDevice_Connect(t *testing.T) {
	if err := iotDev.Connect(); err != nil {
		t.FailNow()
	}
}

func TestIoTDevice_Ping(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := iotDev.Ping(ctx); err != nil {
		t.FailNow()
	}
}

func TestIoTDevice_ObserveInform(t *testing.T) {
	saveFunc := createSaveFunc(t)
	err := iotDev.ObserveInform(saveFunc)
	if err != nil {
		t.FailNow()
	}
}

func TestIoTDevice_IsObserveInformProcess(t *testing.T) {
	isProcess := iotDev.IsObserveInformProcess()
	if !isProcess {
		t.FailNow()
	}
}

func TestIoTDevice_StopObserveInform(t *testing.T) {
	err := iotDev.StopObserveInform()
	if err != nil {
		t.FailNow()
	}
}

func TestIoTDevice_IsObserveInformProcess2(t *testing.T) {
	isProcess := iotDev.IsObserveInformProcess()
	if isProcess {
		t.FailNow()
	}
}

func createSaveFunc(t *testing.T) func([]byte, message.MediaType) error {
	return func(msg []byte, msgType message.MediaType) error {
		t.Log("Got message from save func!")
		return nil
	}
}
