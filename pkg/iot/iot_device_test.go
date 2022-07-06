package iot

import (
	"CoAPProxyServer/pkg/config"
	"context"
	"github.com/plgd-dev/go-coap/v2/message"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"testing"
	"time"
)

var iotDev IoTDevice

func TestInit(t *testing.T) {
	cmd := exec.Command("docker", "build", "../../iotsDevicesImitation/.", "-t", "test_iot")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("unexpected error %s. P.s. for this tests need docker", err)
	}

	cmd = exec.Command("docker", "run",
		"--rm", "-d", "-e", "port=5688", "-e", "inftype=-time", "-p", "5688:5688/udp",
		"--name", "test_iot", "test_iot")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("unexpected error %s. P.s. for this tests need docker", err)
	}
}

func TestIoTDevice_Init(t *testing.T) {
	iotDev = IoTDevice{}
	conf := config.IotConfig{
		Addr: ":5688",
		Name: "testDevice",
	}
	iotDev.Init(conf)

	if iotDev.isObserveInformProcess == nil {
		t.Error("field *bool not initialize")
		return
	}
}

func TestIoTDevice_GetId(t *testing.T) {

}

func TestIoTDevice_GetName(t *testing.T) {
	if iotDev.GetName() != "testDevice" {
		t.Error("unexpected return value")
	}
}

func TestIoTDevice_Connect(t *testing.T) {
	if err := iotDev.Connect(); err != nil {
		t.Errorf("function Connect() is corrupted: unexpected error: %s", err)
	}
}

func TestIoTDevice_Ping(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := iotDev.Ping(ctx); err != nil {
		t.Errorf("function Ping() is corrupted: unexpected error: %s", err)
	}
}

func TestIoTDevice_ObserveInform(t *testing.T) {
	saveFunc := createSaveFunc(t)
	err := iotDev.ObserveInform(saveFunc)
	if err != nil {
		t.Errorf("function ObserveInform() is corrupted: unexpected error: %s", err)
	}
}

func TestIoTDevice_IsObserveInformProcess(t *testing.T) {
	isProcess := iotDev.IsObserveInformProcess()
	if !isProcess {
		t.Errorf("function IsObserveInformProcess() is corrupted: unexpected returned value")
	}
}

func TestIoTDevice_StopObserveInform(t *testing.T) {
	err := iotDev.StopObserveInform()
	if err != nil {
		t.Errorf("function StopObserveInform() is corrupted: unexpected error: %s", err)
	}
}

func TestIoTDevice_IsObserveInformProcess2(t *testing.T) {
	isProcess := iotDev.IsObserveInformProcess()
	if isProcess {
		t.Errorf("function IsObserveInformProcess() is corrupted: unexpected returned value")
	}
}

func TestIoTDevice_Disconnect(t *testing.T) {
	err := iotDev.Disconnect()
	if err != nil {
		t.Errorf("function Disconnect() is corrupted: unexpected error: %s", err)
	}
}

func TestShoutDown(t *testing.T) {
	cmd := exec.Command("docker", "stop", "test_iot")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func createSaveFunc(t *testing.T) func([]byte, message.MediaType) error {
	return func(msg []byte, msgType message.MediaType) error {
		t.Log("Got message from save func!")
		return nil
	}
}
