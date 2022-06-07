package main

import (
	"fmt"
	"github.com/zubairhamed/canopus"
)

func main() {
	conn, err := canopus.Dial("localhost:5683")
	if err != nil {
		panic(err.Error())
	}

	req := canopus.NewRequest(canopus.MessageConfirmable, canopus.Get).(*canopus.CoapRequest)
	req.SetStringPayload("Hello, canopus")
	req.SetRequestURI("/hello")

	resp, err := conn.Send(req)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Got Response:" + resp.GetMessage().GetPayload().String())

}
