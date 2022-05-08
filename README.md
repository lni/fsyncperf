## fsyncperf

This is a very simple (i.e. oversimplified) program to tell you how fast/slow is fsync() on your local disk. 

Go 1.17+ required, to test the fsync() performance of your cwd - 

```
go build .
./fsyncperf
```

## Results

Aliyun ecs.i3.2xlarge's local SSD

```
concurrency: 1
workerID: 0, latency: 182 microsecond per op, bandwidth: 342MBytes/sec
aggregated bandwidth: 342MBytes/sec

concurrency: 2
workerID: 1, latency: 202 microsecond per op, bandwidth: 308MBytes/sec
workerID: 0, latency: 202 microsecond per op, bandwidth: 308MBytes/sec
aggregated bandwidth: 616MBytes/sec

concurrency: 4
workerID: 3, latency: 231 microsecond per op, bandwidth: 269MBytes/sec
workerID: 1, latency: 231 microsecond per op, bandwidth: 269MBytes/sec
workerID: 2, latency: 231 microsecond per op, bandwidth: 269MBytes/sec
workerID: 0, latency: 231 microsecond per op, bandwidth: 269MBytes/sec
aggregated bandwidth: 1076MBytes/sec

concurrency: 8
workerID: 3, latency: 335 microsecond per op, bandwidth: 186MBytes/sec
workerID: 0, latency: 335 microsecond per op, bandwidth: 186MBytes/sec
workerID: 4, latency: 335 microsecond per op, bandwidth: 186MBytes/sec
workerID: 1, latency: 335 microsecond per op, bandwidth: 186MBytes/sec
workerID: 5, latency: 336 microsecond per op, bandwidth: 185MBytes/sec
workerID: 2, latency: 336 microsecond per op, bandwidth: 185MBytes/sec
workerID: 6, latency: 336 microsecond per op, bandwidth: 185MBytes/sec
workerID: 7, latency: 337 microsecond per op, bandwidth: 185MBytes/sec
aggregated bandwidth: 1484MBytes/sec
```

Aliyun ecs.i4p.2xlarge's local SSD (NVRAM based)

```
concurrency: 1
workerID: 0, latency: 102 microsecond per op, bandwidth: 610MBytes/sec
aggregated bandwidth: 610MBytes/sec

concurrency: 2
workerID: 0, latency: 130 microsecond per op, bandwidth: 478MBytes/sec
workerID: 1, latency: 130 microsecond per op, bandwidth: 478MBytes/sec
aggregated bandwidth: 956MBytes/sec

concurrency: 4
workerID: 1, latency: 192 microsecond per op, bandwidth: 324MBytes/sec
workerID: 2, latency: 192 microsecond per op, bandwidth: 324MBytes/sec
workerID: 3, latency: 192 microsecond per op, bandwidth: 324MBytes/sec
workerID: 0, latency: 192 microsecond per op, bandwidth: 324MBytes/sec
aggregated bandwidth: 1296MBytes/sec

concurrency: 8
workerID: 3, latency: 318 microsecond per op, bandwidth: 196MBytes/sec
workerID: 1, latency: 318 microsecond per op, bandwidth: 195MBytes/sec
workerID: 2, latency: 319 microsecond per op, bandwidth: 195MBytes/sec
workerID: 6, latency: 319 microsecond per op, bandwidth: 195MBytes/sec
workerID: 4, latency: 319 microsecond per op, bandwidth: 195MBytes/sec
workerID: 0, latency: 319 microsecond per op, bandwidth: 195MBytes/sec
workerID: 5, latency: 320 microsecond per op, bandwidth: 195MBytes/sec
workerID: 7, latency: 320 microsecond per op, bandwidth: 194MBytes/sec
aggregated bandwidth: 1560MBytes/sec
```

Aliyun ecs.i4p.16xlarge's local SSD (NVRAM based)

```
concurrency: 1
workerID: 0, latency: 71 microsecond per op, bandwidth: 877MBytes/sec
aggregated bandwidth: 877MBytes/sec

concurrency: 2
workerID: 1, latency: 80 microsecond per op, bandwidth: 779MBytes/sec
workerID: 0, latency: 80 microsecond per op, bandwidth: 779MBytes/sec
aggregated bandwidth: 1558MBytes/sec

concurrency: 4
workerID: 3, latency: 96 microsecond per op, bandwidth: 649MBytes/sec
workerID: 0, latency: 96 microsecond per op, bandwidth: 649MBytes/sec
workerID: 1, latency: 96 microsecond per op, bandwidth: 649MBytes/sec
workerID: 2, latency: 96 microsecond per op, bandwidth: 648MBytes/sec
aggregated bandwidth: 2595MBytes/sec

concurrency: 8
workerID: 0, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 7, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 6, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 5, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 1, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 3, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 2, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
workerID: 4, latency: 144 microsecond per op, bandwidth: 432MBytes/sec
aggregated bandwidth: 3456MBytes/sec
```

Samsung PM983A 960GB

```
concurrency: 1
workerID: 0, latency: 299 microsecond per op, bandwidth: 208MBytes/sec
aggregated bandwidth: 208MBytes/sec

concurrency: 2
workerID: 0, latency: 315 microsecond per op, bandwidth: 197MBytes/sec
workerID: 1, latency: 316 microsecond per op, bandwidth: 197MBytes/sec
aggregated bandwidth: 394MBytes/sec

concurrency: 4
workerID: 2, latency: 450 microsecond per op, bandwidth: 138MBytes/sec
workerID: 0, latency: 450 microsecond per op, bandwidth: 138MBytes/sec
workerID: 1, latency: 451 microsecond per op, bandwidth: 138MBytes/sec
workerID: 3, latency: 451 microsecond per op, bandwidth: 138MBytes/sec
aggregated bandwidth: 552MBytes/sec

concurrency: 8
workerID: 0, latency: 631 microsecond per op, bandwidth: 98MBytes/sec
workerID: 5, latency: 632 microsecond per op, bandwidth: 98MBytes/sec
workerID: 2, latency: 632 microsecond per op, bandwidth: 98MBytes/sec
workerID: 6, latency: 633 microsecond per op, bandwidth: 98MBytes/sec
workerID: 3, latency: 633 microsecond per op, bandwidth: 98MBytes/sec
workerID: 7, latency: 632 microsecond per op, bandwidth: 98MBytes/sec
workerID: 4, latency: 633 microsecond per op, bandwidth: 98MBytes/sec
workerID: 1, latency: 634 microsecond per op, bandwidth: 98MBytes/sec
aggregated bandwidth: 784MBytes/sec
```

Intel P4510 2T

```
concurrency: 1
workerID: 0, latency: 151 microsecond per op, bandwidth: 411MBytes/sec
aggregated bandwidth: 411MBytes/sec

concurrency: 2
workerID: 0, latency: 180 microsecond per op, bandwidth: 346MBytes/sec
workerID: 1, latency: 180 microsecond per op, bandwidth: 345MBytes/sec
aggregated bandwidth: 691MBytes/sec

concurrency: 4
workerID: 0, latency: 216 microsecond per op, bandwidth: 288MBytes/sec
workerID: 1, latency: 216 microsecond per op, bandwidth: 288MBytes/sec
workerID: 2, latency: 217 microsecond per op, bandwidth: 287MBytes/sec
workerID: 3, latency: 216 microsecond per op, bandwidth: 288MBytes/sec
aggregated bandwidth: 1151MBytes/sec

concurrency: 8
workerID: 7, latency: 327 microsecond per op, bandwidth: 191MBytes/sec
workerID: 0, latency: 327 microsecond per op, bandwidth: 191MBytes/sec
workerID: 4, latency: 327 microsecond per op, bandwidth: 191MBytes/sec
workerID: 2, latency: 327 microsecond per op, bandwidth: 190MBytes/sec
workerID: 5, latency: 327 microsecond per op, bandwidth: 190MBytes/sec
workerID: 1, latency: 327 microsecond per op, bandwidth: 190MBytes/sec
workerID: 3, latency: 327 microsecond per op, bandwidth: 190MBytes/sec
workerID: 6, latency: 327 microsecond per op, bandwidth: 190MBytes/sec
aggregated bandwidth: 1523MBytes/sec
```

Intel P5510 7.6T

```
concurrency: 1
workerID: 0, latency: 152 microsecond per op, bandwidth: 410MBytes/sec
aggregated bandwidth: 410MBytes/sec

concurrency: 2
workerID: 1, latency: 184 microsecond per op, bandwidth: 339MBytes/sec
workerID: 0, latency: 184 microsecond per op, bandwidth: 339MBytes/sec
aggregated bandwidth: 678MBytes/sec

concurrency: 4
workerID: 2, latency: 249 microsecond per op, bandwidth: 250MBytes/sec
workerID: 3, latency: 249 microsecond per op, bandwidth: 250MBytes/sec
workerID: 1, latency: 250 microsecond per op, bandwidth: 249MBytes/sec
workerID: 0, latency: 252 microsecond per op, bandwidth: 247MBytes/sec
aggregated bandwidth: 996MBytes/sec

concurrency: 8
workerID: 3, latency: 300 microsecond per op, bandwidth: 208MBytes/sec
workerID: 1, latency: 300 microsecond per op, bandwidth: 208MBytes/sec
workerID: 0, latency: 300 microsecond per op, bandwidth: 208MBytes/sec
workerID: 4, latency: 300 microsecond per op, bandwidth: 207MBytes/sec
workerID: 5, latency: 300 microsecond per op, bandwidth: 207MBytes/sec
workerID: 7, latency: 300 microsecond per op, bandwidth: 207MBytes/sec
workerID: 2, latency: 300 microsecond per op, bandwidth: 208MBytes/sec
workerID: 6, latency: 300 microsecond per op, bandwidth: 207MBytes/sec
aggregated bandwidth: 1660MBytes/sec
```


