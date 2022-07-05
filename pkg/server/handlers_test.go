package server

import (
	"CoAPProxyServer/pkg/config"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
	"testing"
)

var (
	proxyServer Server
)

const serverAddr = "127.0.0.1:8080"

type Controller struct {
	ioTs []config.IotConfig
}

func (c *Controller) GetLastNRowsLogs(nRows int) ([]string, error) {
	if nRows < 0 {
		return []string{}, errors.New("wrong count rows")
	}
	return []string{"1 row", "2 row"}, nil
}

func (c *Controller) RemoveIoTDeviceObserve(ioTsConfig []config.IotConfig) error {
	return nil
}

func (c *Controller) NewIotDeviceObserve(iotConfig config.IotConfig) error {
	c.ioTs = append(c.ioTs, iotConfig)
	return nil
}

func (c *Controller) GetInformation(deviceName string) ([]byte, error) {
	for i, t := range c.ioTs {
		if t.Name == deviceName {
			return []byte{byte(i)}, nil
		}
	}

	return []byte{}, errors.New("not found")
}

func Init() {
	cmd := exec.Command("docker", "build", "../../iotsDevicesImitation/.", "-t", "test_iot")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("docker", "run",
		"--rm", "-d", "-e", "port=5688", "-e", "inftype=-time", "--name", "test_iot", "test_iot")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	controller := Controller{}
	go proxyServer.StartServer(config.Config{ProxyServerAddr: serverAddr}, &controller)
}

func ShoutDown() {
	cmd := exec.Command("docker", "stop", "test_iot")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestServer_addIotDevice(t *testing.T) {
	Init() // start only once
	req := httptest.NewRequest(http.MethodGet, "/device/add?deviceName=testName&deviceAddr=:5600", nil)
	w := httptest.NewRecorder()
	proxyServer.addIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestServer_addIotDeviceFail(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/add?deviceAddr=:5600", nil)
	w := httptest.NewRecorder()
	proxyServer.addIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if w.Body.String() == "" {
		t.Fatalf("expected warning msg")
	}
}

func TestServer_addIotDeviceFail2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/add?deviceAddr=:5600", nil)
	w := httptest.NewRecorder()
	proxyServer.addIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if w.Body.String() == "" {
		t.Fatalf("expected warning msg")
	}
}

func TestServer_getInformationFromIotDevice(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/metrics?deviceName=testName", nil)
	w := httptest.NewRecorder()
	proxyServer.getInformationFromIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestServer_getInformationFromIotDeviceFail(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/metrics?", nil)
	w := httptest.NewRecorder()
	proxyServer.getInformationFromIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if w.Body.String() == "" {
		t.Fatalf("expected warning msg")
	}
}

func TestServer_getInformationFromIotDeviceFail2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/metrics?deviceName=wrongName", nil)
	w := httptest.NewRecorder()
	proxyServer.getInformationFromIotDevice(w, req)

	if want, got := http.StatusInternalServerError, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestServer_removeIotDevice(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/rm?deviceName=testName", nil)
	w := httptest.NewRecorder()
	proxyServer.removeIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestServer_removeIotDeviceFail(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/device/rm?", nil)
	w := httptest.NewRecorder()
	proxyServer.removeIotDevice(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if w.Body.String() == "" {
		t.Fatalf("expected warning msg")
	}
}

func TestServer_getLogs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logs?countLogs=2", nil)
	w := httptest.NewRecorder()
	proxyServer.getLogs(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	out := w.Body.String()
	outArr := strings.Split(out, "\n")
	if len(outArr) < 2 {
		t.FailNow()
	}
}

func TestServer_getLogsFail(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logs?", nil)
	w := httptest.NewRecorder()
	proxyServer.getLogs(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if w.Body.String() == "" {
		t.Fatalf("expected warning msg")
	}
}

func TestServer_getLogsFail2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logs?countLogs=q", nil)
	w := httptest.NewRecorder()
	proxyServer.getLogs(w, req)

	if want, got := http.StatusInternalServerError, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestServer_getLogsFail3(t *testing.T) {
	ShoutDown()
	req := httptest.NewRequest(http.MethodGet, "/logs?countLogs=-2", nil)
	w := httptest.NewRecorder()
	proxyServer.getLogs(w, req)

	if want, got := http.StatusInternalServerError, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
