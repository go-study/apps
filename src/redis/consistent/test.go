// Copyright (C) 2012 Numerotron Inc.
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

// Package consistent provides a consistent hashing function.
//
// Consistent hashing is often used to distribute requests to a changing set of servers.  For example,
// say you have some cache servers cacheA, cacheB, and cacheC.  You want to decide which cache server
// to use to look up information on a user.
//
// You could use a typical hash table and hash the user id
// to one of cacheA, cacheB, or cacheC.  But with a typical hash table, if you add or remove a server,
// almost all keys will get remapped to different results, which basically could bring your service
// to a grinding halt while the caches get rebuilt.
//
// With a consistent hash, adding or removing a server drastically reduces the number of keys that
// get remapped.
//
// Read more about consistent hashing on wikipedia:  http://en.wikipedia.org/wiki/Consistent_hashing
//
package main

import (
	"errors"
	"fmt"
	"time"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
	"math/rand"
)

type uints []uint32

// Len returns the length of the uints array.
func (x uints) Len() int { return len(x) }

// Less returns true if element i is less than element j.
func (x uints) Less(i, j int) bool { return x[i] < x[j] }

// Swap exchanges elements i and j.
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Consistent holds the information about the members of the consistent hash circle.
type Consistent struct {
	circle           map[uint32]string
	members          map[string]bool
	sortedHashes     uints // 已经排好序的hashes slice ， 主要有力搜索 (存储的内容是全部虚拟hashes值)
	NumberOfReplicas int
	count            int64
	scratch          [64]byte
	sync.RWMutex
}

// New creates a new Consistent object with a default setting of 20 replicas for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New() *Consistent {
	c := new(Consistent)
	c.NumberOfReplicas = 20
	c.circle = make(map[uint32]string)
	c.members = make(map[string]bool)
	//log.Printf("%p", c)
	return c
}

// eltKey generates a string key for an element with an index.
func (c *Consistent) eltKey(elt string, idx int) string {
	return elt + "|" + strconv.Itoa(idx)
}

// Add inserts a string element in the consistent hash.
func (c *Consistent) Add(elt string) {
	c.Lock()
	defer c.Unlock()
	for i := 0; i < c.NumberOfReplicas; i++ {
		//fmt.Println("i:", i,elt, c.hashKey(c.eltKey(elt, i)))
		c.circle[c.hashKey(c.eltKey(elt, i))] = elt
	}
	c.members[elt] = true

	c.updateSortedHashes()
	c.count++
}

// Remove removes an element from the hash.
func (c *Consistent) Remove(elt string) {
	c.Lock()
	defer c.Unlock()
	for i := 0; i < c.NumberOfReplicas; i++ {
		delete(c.circle, c.hashKey(c.eltKey(elt, i)))
	}
	delete(c.members, elt)
	c.updateSortedHashes()
	c.count--
}

// Set sets all the elements in the hash.  If there are existing elements not present in elts, they will be removed.
func (c *Consistent) Set(elts []string) {
	mems := c.Members()
	for _, k := range mems {
		found := false
		for _, v := range elts {
			if k == v {
				found = true
				break
			}
		}
		if !found {
			c.Remove(k)
		}
	}
	for _, v := range elts {
		c.RLock()
		_, exists := c.members[v]
		c.RUnlock()
		if exists {
			continue
		}
		c.Add(v)
	}
}

func (c *Consistent) Members() []string {
	c.RLock()
	defer c.RUnlock()
	var m []string
	for k := range c.members {
		m = append(m, k)
	}
	return m
}

// Get returns an element close to where name hashes to in the circle.
func (c *Consistent) Get(name string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.circle) == 0 {
		return "", ErrEmptyCircle
	}
	key := c.hashKey(name)
	//log.Println("need search --> key:", key, "servername:", name)
	i := c.search(key)
	//fmt.Println(c.sortedHashes[i], c.circle[c.sortedHashes[i]])
	return c.circle[c.sortedHashes[i]], nil
}

func (c *Consistent) search(key uint32) (i int) {
	f := func(x int) bool {
		//log.Println("i", i)
		// 拿不到相等的
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	//log.Println("I:", i)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (c *Consistent) GetTwo(name string) (string, string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.circle) == 0 {
		return "", "", ErrEmptyCircle
	}
	//得到hashesw 值
	key := c.hashKey(name)
	//搜索hashes
	i := c.search(key)
	//获取值
	a := c.circle[c.sortedHashes[i]]
	//如果节点只有一个时，直接返回
	if c.count == 1 {
		return a, "", nil
	}

	start := i
	var b string
	for i = start + 1; i != start; i++ {
		if i >= len(c.sortedHashes) {
			i = 0
		}
		b = c.circle[c.sortedHashes[i]]
		//两个时候否为相同的节点，不是就返回
		if b != a {
			break
		}
	}
	return a, b, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (c *Consistent) GetN(name string, n int) ([]string, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.circle) == 0 {
		return nil, ErrEmptyCircle
	}

	if c.count < int64(n) {
		n = int(c.count)
	}

	var (
		key   = c.hashKey(name)
		i     = c.search(key)
		start = i
		res   = make([]string, 0, n)
		elem  = c.circle[c.sortedHashes[i]]
	)

	res = append(res, elem)

	if len(res) == n {
		return res, nil
	}

	for i = start + 1; i != start; i++ {
		if i >= len(c.sortedHashes) {
			i = 0
		}
		elem = c.circle[c.sortedHashes[i]]
		if !sliceContainsMember(res, elem) {
			res = append(res, elem)
		}
		if len(res) == n {
			break
		}
	}

	return res, nil
}

func (c *Consistent) hashKey(key string) uint32 {
	//
	//log.Println("key string:", key)
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		//log.Fatal(len(key), scratch)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 对hash 进行排序
func (c *Consistent) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	if cap(c.sortedHashes)/(c.NumberOfReplicas*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
	//log.Println("tem hashes size :", len(hashes), c.sortedHashes)
}

func sliceContainsMember(set []string, member string) bool {
	for _, m := range set {
		if m == member {
			return true
		}
	}
	return false
}

func main() {
	c := New()
	//fmt.Printf("%T", D)
	rand.Seed(time.Now().UnixNano()) 
	c.Add("redis-1")
	c.Add("redis-2")
	c.Add("redis-3")
	c.Add("redis-4")
	c.Add("redis-5")
	c.Add("redis-6")
	fmt.Println(c.Members());
	v,_:= c.Get(string(time.Duration(rand.Intn(5)) * time.Millisecond))
	v,_ =  c.Get("my"); fmt.Println(v);
	v,_ =  c.Get("sx"); fmt.Println(v);
	v,_ =  c.Get("dd1"); fmt.Println(v);
	v,_ =  c.Get("a2"); fmt.Println(v);
	v,_ =  c.Get("g3"); fmt.Println(v);
	v,_ =  c.Get("jfew4"); fmt.Println(v);
	v,_ =  c.Get("ooi5"); fmt.Println(v);
	v,_ =  c.Get("uoi5feiwfowefowejo"); fmt.Println(v);

	//log.Println("members size:", len(c.members), "\tcircle size :", len(c.circle), "sortHashes:", len(c.sortedHashes), "scratch:", c.scratch)
}
