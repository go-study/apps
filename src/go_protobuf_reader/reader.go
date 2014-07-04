package main

import proto "code.google.com/p/goprotobuf/proto"
import (
	"go_protobuf_reader/hello"
	"fmt"
	"io"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func main() {
	path := string("/tmp/log.txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer file.Close()
	fi, err := file.Stat()
	CheckError(err)
	buffer := make([]byte, fi.Size())
	_, err = io.ReadFull(file, buffer) //read all content
	CheckError(err)
	msg := &hello.Hello{}
	err = proto.Unmarshal(buffer, msg) //unSerialize
	CheckError(err)
	fmt.Printf("read: %s\n", msg.String())
}
