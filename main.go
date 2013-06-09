package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/carloscm/gossie/src/gossie"
	"io"
	"os"
	"runtime"
	"strings"
)

var ServerList []string
var ListOpt string
var PoolSizeOpt int
var StepOpt int
var ModeOpt string

func init() {
	flag.StringVar(&ListOpt, "list", "default", "list of server:port to use")
	flag.IntVar(&PoolSizeOpt, "connections", 128, "number of connections to use in the pool")
	flag.IntVar(&StepOpt, "step", 10, "number of samples to average before output")
	flag.StringVar(&ModeOpt, "mode", "read", "which mode to run the tool in (read|write)")

	flag.Parse()

	ServerList = readList(fmt.Sprintf("%s.txt", ListOpt))
}

func readList(path string) (list []string) {
	fd, err := os.Open(path)
	if err != nil {
		fmt.Sprintf("Could not open %s for reading: %s\n", path, err)
		os.Exit(1)
	}
	buf := bufio.NewReader(fd)
	defer fd.Close()

	line, err := buf.ReadString('\n')
	for err != io.EOF {
		list = append(list, strings.Trim(line, "\n"))
		line, err = buf.ReadString('\n')
	}

	return list
}

func main() {
	runtime.GOMAXPROCS(11)

	pool, err := gossie.NewConnectionPool(ServerList, "gostress", gossie.PoolOptions{Size: PoolSizeOpt, Timeout: 30000})
	if err != nil {
		fmt.Sprintf("Error connecting to cluster: %v\n", err)
		return
	}

	switch ModeOpt {
	case "read":
		readLoad(pool, ServerList)
	case "write":
		writeLoadData(pool, ServerList)
	default:
		{
			fmt.Printf("Invalid -mode value: '%s'\n", ModeOpt)
		}
	}
}
