// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net"
	"time"
	"net/http"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)


  func newPool(server, password string) *redis.Pool {
      return &redis.Pool{
          MaxIdle: 3,
          IdleTimeout: 240 * time.Second,
          Dial: func () (redis.Conn, error) {
              c, err := redis.Dial("tcp", server)
              if err != nil {
                  return nil, err
              }
              if _, err := c.Do("AUTH", password); err != nil {
                  c.Close()
                  return nil, err
              }
              return c, err
          },
          TestOnBorrow: func(c redis.Conn, t time.Time) error {
              _, err := c.Do("PING")
              return err
          },
      }
  }

  var (
      pool *redis.Pool
      redisServer = flag.String("redisServer", ":6379", "")
      redisPassword = flag.String("redisPassword", "", "")
  )
func redisGetHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	key := r.FormValue("key")

	conn := pool.Get()
	defer conn.Close()
	s,err := conn.Do("GET",key);



	if err != nil {
		fmt.Fprintf(w, "/values: is empty%s/", s)
		return
	}

	fmt.Fprintf(w, "/values:%s/", s)
}
func redisSetHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	value := r.FormValue("value")
	key := r.FormValue("key")
	fmt.Fprintf(w, "/set key:%s,value:%s/", key, value)

	conn := pool.Get()
	defer conn.Close()
	_,err = conn.Do("SET",key,value);

	if err != nil {
		fmt.Fprintf(w, "/set values: is empty%s/")
		return
	}

	fmt.Fprintf(w, "/set key:%s,value:%s/", key, value)
}
func main() {
	flag.Parse()
		pool = newPool(*redisServer, *redisPassword)
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
