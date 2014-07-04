package main

import proto "code.google.com/p/goprotobuf/proto"
import (
	"go_protobuf/hello"
	"fmt"
	"os"
	"bytes"
)
import "net/http"
import "io/ioutil"

func main() {

	
	imgUrl :=&hello.ImgUrl{Url:proto.String("http://test")}
	msg := &hello.User{
		Uid:  proto.Int32(101),
		Uname: proto.String("hello"),
		Imgurl: []*hello.ImgUrl{imgUrl},
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

	client := &http.Client{};
	req, _ := http.NewRequest("POST", "http://112.124.46.89:94/protocolbuf/test/test.php", bytes.NewBuffer(buffer))
		req.Header.Add("Content-type","application/octet-stream");
	resp, _ := client.Do(req);
	defer resp.Body.Close()
	fmt.Println("HTTP returned status %v", resp.Status)
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("HTTP returned status %v", string(bodyByte))
}
