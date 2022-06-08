package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/plgd-dev/go-coap/v2/udp"
	"github.com/plgd-dev/go-coap/v2/udp/message/pool"
)

func main() {
	sync := make(chan bool)

	co, err := udp.Dial("localhost:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	num := 0
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*9)
	defer cancel()

	resp, err := co.Get(ctx, "/time")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.String())

	obs, err := co.Observe(ctx, "/some/path", func(req *pool.Message) {
		log.Printf("Got %+v\n", req)
		buff := make([]byte, 300)
		_, err := req.Body().Read(buff)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buff))
		num++
		if num >= 10 {
			sync <- true
		}
	})
	if err != nil {
		log.Fatalf("Unexpected error '%v'", err)
	}

	<-sync
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	obs.Cancel(ctx)
}
