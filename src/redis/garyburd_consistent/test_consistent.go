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
	"net/http"
	"redis/consistent2"
	"sync"
	"time"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
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

func redisGetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	s, ok := hash[key]
	if ok {
		fmt.Fprintf(w, "/111values:%s/", s)
		return
	} else {
		conn := getPool(key).Get()
		defer conn.Close()
		s, err := conn.Do("GET", key)
		if err != nil {
			fmt.Fprintf(w, "/values3333: is empty%s/", s)
		}
		fmt.Fprintf(w, "/2222222values:%s/", s)
		return
	}

}
func redisSetHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	value := r.FormValue("value")
	key := r.FormValue("key")
	fmt.Fprintf(w, "/set key:%s,value:%s/", key, value)

	conn := getPool(key).Get()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	//
	done.Lock()
	hash[key] = value
	done.Unlock()

	if err != nil {
		fmt.Fprintf(w, "/set values: is empty%s/")
		return
	}

	fmt.Fprintf(w, "/set key:%s,value:%s/", key, value)
}

var (
	pool          *redis.Pool
	redisServer   = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "", "")
)
var (
	pools = make(map[string]*redis.Pool)
	c     = Consistent.New()
	hash  = make(map[string]string)
	done  sync.RWMutex
)

func getPool(key string) *redis.Pool {
	key_index, _ := c.Get(key)
	return pools[key_index]
}
func main() {

	//var c = Consistent.New()
	c.Add("redis-1")
	c.Add("redis-2")
	c.Add("redis-3")
	//c.Add("redis-4")
	//c.Add("redis-5")
	//c.Add("redis-6")

	//flag.Parse()
	//pool = newPool(*redisServer, *redisPassword)
	pools["redis-1"] = newPool(*redisServer, *redisPassword)
	pools["redis-2"] = newPool(*redisServer, *redisPassword)
	pools["redis-3"] = newPool(*redisServer, *redisPassword)

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
