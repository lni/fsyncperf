package main

import (
	"flag"
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

func syncStart(workerID uint64, concurrency uint64, ch chan struct{}) {
	return
	if workerID == 0 {
		for i := uint64(1); i < concurrency; i++ {
			ch <- struct{}{}
		}
	} else {
		<-ch
	}
}

func syncEnd(workerID uint64, concurrency uint64, ch chan struct{}) {
	return
	if workerID == 0 {
		for i := uint64(1); i < concurrency; i++ {
			<-ch
		}
	} else {
		ch <- struct{}{}
	}
}

func writeFsyncTest(workerID uint64, ch chan result,
	syncStartCh chan struct{}, syncEndCh chan struct{}, concurrency uint64, path string) {
	p := path + "/" + dataFilename
	fn := fmt.Sprintf(p, workerID)
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
		syncStart(workerID, concurrency, syncStartCh)
		if _, err := f.Write(buf); err != nil {
			ch <- result{err: err}
			return
		}
		if err := f.Sync(); err != nil {
			ch <- result{err: err}
			return
		}
		syncStart(workerID, concurrency, syncEndCh)
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

func test(concurrency uint64, path string) {
	resultCh := make(chan result, concurrency)
	syncStartCh := make(chan struct{})
	syncEndCh := make(chan struct{})
	for workerID := uint64(0); workerID < concurrency; workerID++ {
		go writeFsyncTest(workerID, resultCh, syncStartCh, syncEndCh, concurrency, path)
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
	path := flag.String("path", "path", "ebs device name list")
	flag.Parse()

	fmt.Printf("DeviceName: %s \n", *path)
	test(1, *path)
	test(2, *path)
	test(4, *path)
	test(8, *path)
}
