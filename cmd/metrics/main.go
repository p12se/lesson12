package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	graphite "github.com/cyberdelia/go-metrics-graphite"
	"github.com/rcrowley/go-metrics"
)

const fanout = 10

func main() {
	prefix := os.Args[1]

	c := metrics.NewCounter()
	metrics.Register("foo", c)

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:2003")
	go graphite.Graphite(metrics.DefaultRegistry, 1, prefix, addr)

	var a int64 = 47
	for range time.Tick(time.Second * 30) {
		a = a + rand.Int63n(100)
		c.Inc(a)
		fmt.Println(a)
	}

}
