## fsyncperf

This is a very simple (i.e. oversimplified) program to tell you how fast/slow is fsync() on your local disk. 

Go 1.17+ required, to test the fsync() performance of your cwd - 

```
go build .
./fsyncperf
```
