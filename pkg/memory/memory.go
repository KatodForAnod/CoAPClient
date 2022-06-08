package memory

import "github.com/plgd-dev/go-coap/v2/message"

type Memory interface {
	Save([]byte, message.MediaType) error
}
