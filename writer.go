package main

import (
	"crypto/rand"
	"fmt"
	"github.com/carloscm/gossie/src/gossie"
)

func generateColumns(prefix string, count int64) []*gossie.Column {
	columns := []*gossie.Column{}

	fmt.Printf("Creating %s test columns...", count)
	var i int64
	for i = 0; i < count; i++ {
		buf := make([]byte, 8192, 8192)
		_, err := rand.Read(buf)
		if err != nil {
			fmt.Printf("rand.Read failed: %s\n", err)
		}
		ck := fmt.Sprintf("col-%s-%d", prefix, i)
		col := gossie.Column{[]byte(ck), buf, 0, -1}
		columns = append(columns, &col)
	}

	return columns
}

func writeLoadData(pool gossie.ConnectionPool, servers []string) {
	for _, server := range servers {
		fmt.Printf("%s - creating row\n", server)
		columns := generateColumns(server, 1e3)
		row := gossie.Row{[]byte(server), columns}
		fmt.Printf("%s - created row\n", server)

		err := pool.Writer().Insert("stressful", &row).Run()

		if err != nil {
			fmt.Printf("Error inserting row on %s: %v\n", server, err)
		}
		fmt.Printf("%s done\n", server)
	}
}
