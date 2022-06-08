package main

import (
	"CoAPClient/pkg/config"
	"github.com/plgd-dev/go-coap/v2/udp"
	"github.com/plgd-dev/go-coap/v2/udp/client"
	"github.com/plgd-dev/go-coap/v2/udp/message/pool"
	"golang.org/x/net/context"
	"log"
)

type IoTDevice struct {
	addr string
	name string
	conn *client.ClientConn
}

func (d *IoTDevice) Init(config config.IotConfig) {
	d.addr = config.Addr
	d.name = config.Name
}

func (d *IoTDevice) Ping(ctx context.Context) error {
	log.Println("ping iot", d.name, "device")
	if err := d.conn.Ping(ctx); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *IoTDevice) ObserveInform(ctx context.Context,
	processMsg func(req *pool.Message)) error {
	log.Println("observe information iot", d.name)
	observe, err := d.conn.Observe(ctx, "/some/path", processMsg)
	if err != nil {
		log.Println(observe)
		return err
	}

	return nil
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
