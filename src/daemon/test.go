package main
import (
	"util"
	"fmt"
	"time"
	)
      
func main(){
  util.Daemon(0,1);
  for{
      fmt.Println("hello")
              time.Sleep(1 * time.Second)
  }
}
