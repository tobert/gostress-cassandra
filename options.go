package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var ServerList *[]string
var PoolSizeOpt int
var StepOpt int
var ModeOpt string

func init() {
	var listFlag string

	flag.StringVar(&listFlag, "list", "default", "list of server:port to use")
	flag.IntVar(&PoolSizeOpt, "connections", 128, "number of connections to use in the pool")
	flag.IntVar(&StepOpt, "step", 10, "number of samples to average before output")
	flag.StringVar(&ModeOpt, "mode", "read", "which mode to run the tool in (read|write)")

	flag.Parse()

	list := readList(fmt.Sprintf("%s.txt", listFlag))
	ServerList = &list
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
