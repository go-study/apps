package main
 
import "fmt"
import "time"
import "sync/atomic"
 
func main() {
    var cnt uint32 = 0
    for i := 0; i < 10; i++ {
        go func() {
            for i:=0; i<20; i++ {
                time.Sleep(time.Millisecond)
                atomic.AddUint32(&cnt, 1)
            }
        }()
    }
    time.Sleep(time.Second)//等一秒钟等goroutine完成
    cntFinal := atomic.LoadUint32(&cnt)//取数据
    fmt.Println("cnt:", cntFinal)
}
