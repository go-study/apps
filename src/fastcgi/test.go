package main
/*
server {
        listen 80;
        server_name go.dev;
        root /root/go/src/godev;
        index index.html;
        #gzip off;
        #proxy_buffering off;

        location / {
                 try_files $uri $uri/;
        }

        location ~ /app.* {
                include         fastcgi.conf;
                fastcgi_pass    127.0.0.1:9001;
        }

        try_files $uri $uri.html =404;
}
//http://go.dev/app
*/

import (
	"net"
	"net/http"
	"net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
