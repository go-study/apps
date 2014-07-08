package main

import (
        "image/png"
)
import "github.com/qpliu/qrencode-go/qrencode"
import _ "fmt"
import _ "strings"
import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

var err error
var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	grid, err := qrencode.Encode("Testing one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixtee n seventeen eighteen nineteen twenty.", qrencode.ECLevelQ)
        if err != nil {
                return
        }
        png.Encode(w, grid.Image(8))
}


func main() {
	flag.Parse()
	http.HandleFunc("/", makeHandler(viewHandler))
	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}
