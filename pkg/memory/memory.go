package memory

import (
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
)

type Memory interface {
	Save([]byte, message.MediaType) error
}

type MemoryFmt struct{}

func (f MemoryFmt) Save(msg []byte, typeMsg message.MediaType) error {
	fmt.Println(string(msg), typeMsg.String())
	return nil
}
