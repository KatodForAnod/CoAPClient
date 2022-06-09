package main

import (
	"log"
	"os"
)

func logInit() error {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	return nil
}

func OpenLastLogFile() (*os.File, error) {
	return nil, nil
}

func ReadNLastRows(file *os.File) ([]string, error) {
	return nil, nil
}
