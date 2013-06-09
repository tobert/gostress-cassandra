package main

import (
	"fmt"
	"runtime"
	"github.com/carloscm/gossie/src/gossie"
)

func main() {
	runtime.GOMAXPROCS(11)

	pool, err := gossie.NewConnectionPool(*ServerList, "gostress", gossie.PoolOptions{Size: PoolSizeOpt, Timeout: 30000})
	if err != nil {
		fmt.Sprintf("Error connecting to cluster: %v\n", err)
		return
	}

	switch ModeOpt {
		case "read":  readLoad(pool)
		case "write": writeLoadData(pool)
		default: {
			fmt.Printf("Invalid -mode value: '%s'\n", ModeOpt)
		}
	}
}

