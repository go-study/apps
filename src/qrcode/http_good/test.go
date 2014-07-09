package main

import (
         "image/png"
         "image"
	_ "strconv"
)
import "code.google.com/p/rsc/qr"
import _ "code.google.com/p/rsc/qr/coding"
import "code.google.com/p/rsc/qr/web/resize"
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
	c, err := qr.Encode(r.FormValue("t"), qr.L)
        if err != nil {
        }
        pngdat := c.Image()
	//fmt.Println(pngdat.Bounds());
	//fmt.Println(c.Size);
	//fmt.Println(c.Scale);
	//to do resize
	newImage := resize.Resample(pngdat, image.Rect(0,0,c.Size,c.Size),120,120);
        png.Encode(w,newImage);
        //png.Encode(w,pngdat);
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

	http.ListenAndServe(":8081", nil)
}
