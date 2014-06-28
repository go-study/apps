package main

import (
	"fmt"
	"time"
	"util"
)

func main() {
	util.Daemon(0, 1)
	for {
		fmt.Println("hello")
		time.Sleep(1 * time.Second)
	}
}
