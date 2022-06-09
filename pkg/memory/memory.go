package memory

import (
	"bufio"
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
	"io/ioutil"
	"log"
	"os"
)

type Memory interface {
	Save([]byte, message.MediaType) error
	Load() ([]byte, error)
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
	log.Println("init membuff")
	b.fileName = fileName
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	writer := bufio.NewWriter(file)

	b.writer = writer
	return nil
}

func (b *MemBuff) Save(msg []byte, typeMsg message.MediaType) error {
	log.Println("save in membuff")
	_, err := b.writer.WriteString(string(msg[:len(msg)]) + "\n")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (b *MemBuff) Load() ([]byte, error) {
	log.Println("load from membuff")
	file, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (b *MemBuff) FlushToFile() error {
	log.Println("fluash to file in membuff")
	if err := b.writer.Flush(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
