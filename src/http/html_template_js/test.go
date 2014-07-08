package main
import "bytes"
import _ "fmt"
import _ "strings"
import (
	"flag"
	"html/template"
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
//var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)\\.js$")

func makeHandler(fn func(http.ResponseWriter, *http.Request )) http.HandlerFunc {
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
	//fmt.Println("D");
	data:=map[string]string{}
	data["js"]="het";
	//t,err = template.ParseFiles("main.orgin.js")
	//if(err !=nil){
	//	panic(err);
	//}
if(len(o)<=0){
	var out bytes.Buffer;
	err = t.Execute(&out,data);
	//if(err !=nil){
	//	panic(err);
	//}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	o=out.Bytes();
}
	//fmt.Printf("%s",out.String());
//	fmt.Println(len(o));
	w.Write(o);
}
var t * template.Template
var o []byte;
func main(){
	t,err = template.ParseFiles("main.orgin.js")
	if(err !=nil){
		panic(err);
	}
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
