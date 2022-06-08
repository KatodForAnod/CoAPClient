package iot

import (
	"CoAPClient/pkg/config"
	"errors"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/udp"
	"github.com/plgd-dev/go-coap/v2/udp/client"
	"github.com/plgd-dev/go-coap/v2/udp/message/pool"
	"golang.org/x/net/context"
	"log"
	"time"
)

type IoTDevice struct {
	addr string
	name string
	conn *client.ClientConn

	observe                *client.Observation
	isObserveInformProcess *bool
}

func (d *IoTDevice) Init(config config.IotConfig) {
	d.addr = config.Addr
	d.name = config.Name

	d.isObserveInformProcess = new(bool)
}

func (d *IoTDevice) Ping(ctx context.Context) error {
	log.Println("ping iot", d.name, "device")
	if d.conn == nil {
		err := errors.New("nil connection if iot " + d.name)
		log.Println(err)
		return err
	}

	if err := d.conn.Ping(ctx); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *IoTDevice) ObserveInform(ctx context.Context,
	save func([]byte, message.MediaType) error) error {
	log.Println("observe information iot", d.name)

	sync := make(chan bool)
	processMsg := func(req *pool.Message) {
		log.Printf("Got %+v\n", req)
		buff := make([]byte, 300)
		if _, err := req.Body().Read(buff); err != nil {
			log.Println(err)
			sync <- true
			return
		}
		infType, err := req.Message.ContentFormat()
		if err != nil {
			log.Println(err)
			sync <- true
			return
		}
		if err := save(buff, infType); err != nil {
			log.Println(err)
			sync <- true
			return
		}
	}

	b := true
	d.isObserveInformProcess = &b
	observe, err := d.conn.Observe(ctx, "/some/path", processMsg)
	if err != nil {
		log.Println(observe)
		b := false //check change
		d.isObserveInformProcess = &b
		return err
	}

	<-sync
	d.observe = observe
	if err := d.StopObserveInform(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *IoTDevice) StopObserveInform() error {
	log.Println("stop observe information")
	b := false
	d.isObserveInformProcess = &b
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := d.observe.Cancel(ctx); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *IoTDevice) IsObserveInformProcess() bool {
	return *d.isObserveInformProcess
}

func (d *IoTDevice) Connect() error {
	log.Println("connecting to iot", d.name)
	conn, err := udp.Dial(d.addr)
	if err != nil {
		log.Printf("Error dialing: %v\n", err)
		return err
	}

	d.conn = conn
	return nil
}

func (d *IoTDevice) Disconnect() error {
	log.Println("disconnecting from iot", d.name)
	err := d.conn.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
