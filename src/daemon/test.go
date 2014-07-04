package main

import (
	"fmt"
	"time"
	"daemon/lib"
)

func main() {
	lib.Daemon(0, 1)
	for {
		fmt.Println("hello")
		time.Sleep(1 * time.Second)
	}
}
