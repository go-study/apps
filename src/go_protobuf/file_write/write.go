package main

import proto "code.google.com/p/goprotobuf/proto"
import (
	"go_protobuf/hello"
	"fmt"
	"os"
)

func main() {

	msg := &hello.Hello{
		Id:  proto.Int32(101),
		Str: proto.String("hello"),
	} //msg init

	path := string("/tmp/log.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer f.Close()
	buffer, err := proto.Marshal(msg) //SerializeToOstream
	f.Write(buffer)
}
