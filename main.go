package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	concurrency  = 8
	dataFilename = "fsync_perf_data_%d.tmp"
	blockSize    = 64 * 1024
	iteration    = 20000
)

type result struct {
	err       error
	workerID  uint64
	latency   int64
	bandwidth int64
}

func writeFsyncTest(workerID uint64, ch chan result) {
	fn := fmt.Sprintf(dataFilename, workerID)
	f, err := os.Create(fn)
	if err != nil {
		ch <- result{err: err}
		return
	}
	defer os.RemoveAll(fn)
	defer f.Close()

	buf := make([]byte, blockSize)
	rand.Read(buf)

	st := time.Now().UnixMicro()
	for i := 0; i < iteration; i++ {
		if _, err := f.Write(buf); err != nil {
			ch <- result{err: err}
			return
		}
		if err := f.Sync(); err != nil {
			ch <- result{err: err}
			return
		}
	}
	total := time.Now().UnixMicro() - st

	ch <- result{
		workerID:  workerID,
		latency:   total / iteration,
		bandwidth: int64(blockSize*iteration) * 1000000 / (total * 1024 * 1024),
	}
}

func print(results []result) {
	fmt.Printf("concurrency: %d\n", len(results))
	bandwidth := int64(0)
	for _, rec := range results {
		bandwidth += rec.bandwidth
		fmt.Printf("workerID: %d, latency: %d microsecond per op, bandwidth: %dMBytes/sec\n",
			rec.workerID, rec.latency, rec.bandwidth)
	}
	fmt.Printf("aggregated bandwidth: %dMBytes/sec\n", bandwidth)
	fmt.Printf("\n")
}

func test(concurrency uint64) {
	resultCh := make(chan result, concurrency)
	for workerID := uint64(0); workerID < concurrency; workerID++ {
		go writeFsyncTest(workerID, resultCh)
	}

	completed := uint64(0)
	results := make([]result, 0)
	for {
		result := <-resultCh
		if result.err != nil {
			panic(result.err)
		}
		results = append(results, result)
		completed++
		if completed == concurrency {
			print(results)
			return
		}
	}
}

func main() {
	test(1)
	test(2)
	test(4)
	test(8)
}
