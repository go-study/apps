package main

import "fmt"
import "time"
import "math/rand"
import "runtime"
import "sync"

var total_tickets int32 = 10

var mutex sync.Mutex
func sell_tickets(i int) {
	for {
		mutex.Lock()
			if total_tickets >0 { //如果有票就卖
				time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
					total_tickets-- //卖一张票
					fmt.Println("id:", i, "  ticket:", total_tickets)
			} else {
				break
			}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().Unix()) //生成随机种子

	for i := 0; i < 5; i++ { //并发5个goroutine来卖票
		go sell_tickets(i)
	}
	//等待线程执行完
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done") //退出时打印还有多少票
}
