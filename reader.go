package main

import (
	"fmt"
	"github.com/carloscm/gossie/src/gossie"
	"time"
)

func readLoader(server string, pool gossie.ConnectionPool, out chan string, died chan bool) {
	var tick int64 = 1
	var totalLatency float64 = 0
	var bytesRead int = 0
	var rowKey []byte = []byte(server)

	for {
		start := time.Now().UnixNano()
		data, err := pool.Reader().Cf("stressful").Get(rowKey)
		done := time.Now().UnixNano()

		if err != nil {
			fmt.Printf("Error? %s -> %v\n", server, err)
			time.Sleep(time.Second * 1)
			continue
		} else {
			for _, val := range data.Columns {
				// ttl + timestamp is 12 bytes, this is just an approximation
				bytesRead = bytesRead + len(val.Name) + len(val.Value) + 12
			}
			totalLatency = totalLatency + float64(done-start)
		}

		if tick%int64(StepOpt) == 0 {
			out <- fmt.Sprintf("%s,%d,%f,%d\n", server, tick/int64(StepOpt), totalLatency/100000000, bytesRead)
			totalLatency = 0
			bytesRead = 0
		}

		tick = tick + 1
	}
	died <- true
}

func readLoad(pool gossie.ConnectionPool, servers []string) {

	out := make(chan string)
	died := make(chan bool)

	go func() {
		for output := range out {
			fmt.Print(output)
		}
		died <- true
	}()

	// CSV header
	fmt.Printf("server,count,latency,bytes\n")

	// one reader per server in the cluster seems to be parallel enough
	for _, server := range servers {
		go readLoader(server, pool, out, died)
	}

	// block until >= 1 goroutines exits
	<-died
}
