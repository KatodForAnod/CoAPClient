package memory

import (
	"fmt"
	"github.com/plgd-dev/go-coap/v2/message"
	log "github.com/sirupsen/logrus"
	"os"
)

type Memory interface {
	Save(msg []byte, typeMsg message.MediaType, nameDevice string) error
	Load(nameDevice string) ([]byte, error)
}

type MemoryFmt struct{}

func (f MemoryFmt) Save(msg []byte, typeMsg message.MediaType, nameDevice string) error {
	fmt.Println(string(msg), typeMsg.String())
	return nil
}

func (f MemoryFmt) Load(nameDevice string) ([]byte, error) {
	return nil, nil
}

type MemBuff struct {
	buffers map[string][]byte
}

func (b *MemBuff) InitStruct() error {
	log.Println("init membuff")
	/*b.fileName = fileName
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	writer := bufio.NewWriter(file)*/
	b.buffers = make(map[string][]byte)

	return nil
}

func (b *MemBuff) Save(msg []byte, typeMsg message.MediaType, nameDevice string) error {
	buff, isExist := b.buffers[nameDevice]
	if !isExist {
		newBuff := make([]byte, 0)
		buff = append(newBuff, msg...)
	}
	buff = append(buff, msg...)
	b.buffers[nameDevice] = buff

	return nil
}

func (b *MemBuff) Load(nameDevice string) ([]byte, error) {
	buff, isExist := b.buffers[nameDevice]
	if !isExist {
		return []byte{}, fmt.Errorf("device %s not found", nameDevice)
	}

	return buff, nil
}

func (b *MemBuff) FlushToFile(nameDevice string) error {
	log.Println("fluash to file in membuff")
	file, err := os.Create(nameDevice + ".txt") // if already exist??
	if err != nil {
		return fmt.Errorf("flust to file err: %v", err)
	}

	buff, isExist := b.buffers[nameDevice]
	if !isExist {
		return fmt.Errorf("device %s not found", nameDevice)
	}
	_, err = file.Write(buff)
	if err != nil {
		return fmt.Errorf("file write err: %v", err)
	}
	b.buffers[nameDevice] = []byte{}

	return nil
}
