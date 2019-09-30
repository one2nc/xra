package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const JUST_HELLO = 0
const JUST_HELLO_RESPONSE = 1
const JUST_BYE = 2
const PERF_START = 3
const PERF_END = 4
const PERF_PAYLOAD = 5

const PERF_MESSAGE_LENGTH = 1024 * 1024 * 1024

type PerfMessageHeader struct {
	Type      int
	Length    int
	TimeStamp time.Time
}

func GetPerfMessageHeader(msgtype int, length int) PerfMessageHeader {
	return PerfMessageHeader{Type: msgtype, Length: length, TimeStamp: time.Now()}
}

type Message []byte

func createRandomMessage(size int) (message Message, err error) {
	fd, err := os.Open("/dev/urandom")

	if err != nil {
		errmsg := fmt.Sprintf("\nError while reading file")
		return nil, errors.New(errmsg)
	}
	defer fd.Close()

	buffer := make([]byte, size)

	fd.Read(buffer)

	return
}
