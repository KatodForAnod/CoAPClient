package memory

import (
	"bufio"
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
	"log"
	"os"
)

type Memory interface {
	Save([]byte, message.MediaType) error
}

type MemoryFmt struct{}

func (f MemoryFmt) Save(msg []byte, typeMsg message.MediaType) error {
	fmt.Println(string(msg), typeMsg.String())
	return nil
}

type MemBuff struct {
	fileName string
	writer   *bufio.Writer
}

func (b *MemBuff) InitStruct(fileName string) error {
	b.fileName = fileName
	file, err := os.Create(fileName)
	writer := bufio.NewWriter(file)
	if err != nil {
		log.Println(err)
		return err
	}

	b.writer = writer
	return nil
}

func (b *MemBuff) Save(msg []byte, typeMsg message.MediaType) error {
	_, err := b.writer.WriteString(string(msg[:len(msg)]) + "\n")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *MemBuff) FlushToFile() error {
	if err := b.writer.Flush(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
