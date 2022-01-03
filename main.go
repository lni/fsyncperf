package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	dataFilename = "fsync_perf_data.tmp"
	iteration    = 100
)

func main() {
	f, err := os.Create(dataFilename)
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dataFilename)
	defer f.Close()

	buf := make([]byte, 1024*64)
	rand.Read(buf)

	st := time.Now().UnixMilli()
	for i := 0; i < iteration; i++ {
		if _, err := f.Write(buf); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
	}
	total := time.Now().UnixMilli() - st
	// yes, latency might be 0 here on a fast disk, you end up with a div by 0
	// crash, good luck for that
	latency := total / iteration
	ops := 1000 / latency
	fmt.Printf("latency: %d millisecond per fsync, %d ops\n", latency, ops)
}
