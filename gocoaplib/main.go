package main

/*
import (
	"bytes"
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/udp"
	"log"
	"os"
	"time"
)


import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/plgd-dev/go-coap/v2/udp"
)

func main22() {
	co, err := udp.Dial("localhost:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	path := "/b"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := co.Get(ctx, path)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	a := make([]byte, 300)
	resp.Body().Read(a)
	log.Printf("Response payload: %v", resp.String())
	fmt.Println(string(a))

	resp, err := co.Post(ctx, path, message.TextPlain, bytes.NewReader([]byte("B hello world")))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	log.Printf("Response payload: %v", resp.String())
}
*/
