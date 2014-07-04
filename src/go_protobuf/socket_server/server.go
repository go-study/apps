package main
import (
    "fmt"
    "os"
    "net"
    "code.google.com/p/goprotobuf/proto"
    "go_protobuf/hello"
)
func connHandle(conn net.Conn) {
    defer conn.Close()
    buff := make([]byte, 128)
    ReadLen, err := conn.Read(buff)
    if err != nil {
        fmt.Println("read data failed...")
        os.Exit(1)
    }
    fmt.Printf("read len: %d\n", ReadLen)
    fmt.Println(buff)
    //根据接收到的数据长度 取出真正的数据  否则比较是否符合规范会出错
    MsgBuf := buff[0 : ReadLen]
    ReciveMsg := &hello.User{}
    //反序列化数据 并判断数据是否符合protobuf协议规范
    err = proto.Unmarshal(MsgBuf, ReciveMsg)
    if err != nil {
        fmt.Printf("unmarshaling error: ", ReciveMsg)
    }
    fmt.Printf("msg id: %d\n", ReciveMsg.GetUid())
    fmt.Printf("msg info: %s\n", ReciveMsg.GetUname())
}
func main() {
    //获取TCP地址
    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":2121")
    if err != nil {
        fmt.Println("get tcp addr failed...")
        os.Exit(1)
    }
    //开始监听
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        fmt.Println("listen tcp failed...")
        os.Exit(1)
    }
    //循环监听
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        //子goroutine中处理
        go connHandle(conn)
    }
}
