# HTTP Perf

Measuring performance of http endpoints.

## Usage

Currently measures for an amount of time (e.g. default 60s) or for a specific count of requests.
Creates a Graph, outputs the avg ms per request and the amount of requests made.

```
http-perf -url http://localhost:8080/
```

## Options
```
▶ http-perf -h
Usage of http-perf:
  -chartpath string
        path for chart png (default "perf.png")
  -count int
        how many times shoud be measured. If used, time flag will be ignored
  -n int
        amount of goroutines beeing used (default 100)
  -nochart
        set if no chart should be generated
  -time int
        how many seconds should be measured. Will be ignored if amount flag is set (default 60)
  -url string
        url which should be measured
```

## Example

```
▶ http-perf -url http://localhost:8080 -time 20 -n 1
2017/08/05 19:02:51 Measuring for 20 seconds with 1 goroutines
2017/08/05 19:03:11 Do not start another request
2017/08/05 19:03:11 Took 20.003788 seconds for 3607 measurements
2017/08/05 19:03:11 AVG request time: 5.533430 ms
2017/08/05 19:03:11 Creating Chart...
2017/08/05 19:03:12 Created Chart sucessfully at perf.png
```

![](./perf_example.png)