// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/gosexy/redis"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

func redisGetHandler(w http.ResponseWriter, r *http.Request) {
	var client *redis.Client
	var err error
	client = redis.New()
	err = client.Connect("localhost", 6379)
	if err != nil {
		fmt.Fprintf(w, "/connect fail/")
		return
	}
	defer client.Quit()

	key := r.FormValue("key")

	s, err := client.Get(key)

	if err != nil {
		fmt.Fprintf(w, "/values: is empty%s/", s)
		return
	}

	fmt.Fprintf(w, "/values:%s/", s)
}
func redisSetHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var client *redis.Client
	client = redis.New()
	err = client.Connect("localhost", 6379)
	if err != nil {
		fmt.Fprintf(w, "/connect fail/")
		return
	}
	defer client.Quit()
	value := r.FormValue("value")
	key := r.FormValue("key")

	client.Set(key, value)

	if err != nil {
		fmt.Fprintf(w, "/set values: is empty%s/")
		return
	}

	fmt.Fprintf(w, "/set key:%s,value:%s/", key, value)
}
func main() {
	flag.Parse()
	http.HandleFunc("/redis/get/", redisGetHandler)
	http.HandleFunc("/redis/set/", redisSetHandler)

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
