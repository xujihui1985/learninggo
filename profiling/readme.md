### Types of Profiling

CPU Profiling

```bash
GODEBUG=schedtrace=1000 ./webapp > /dev/null
print trace every 1 second, as trace will print to stderr, so we redirect stdout to /dev/null
```

```
SCHED 0ms: gomaxprocs=4 idleprocs=2 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0]
SCHED 1002ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 2012ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 3022ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 4032ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 5036ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 6039ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 7050ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 8060ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 9070ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
SCHED 10076ms: gomaxprocs=4 idleprocs=4 threads=10 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
```

the output print every second, total go proc 4, total idle proc 4, machine level thread 10, global runqueue 0 and local runqueue is all 0

we can put load on our server by using hey

```
hey -m POST -c 100 -n 100000 "http://localhost:5000/search
```

now the load is 

```
SCHED 16088ms: gomaxprocs=4 idleprocs=0 threads=15 spinningthreads=1 idlethreads=2 runqueue=6 [0 0 0 0]
SCHED 17095ms: gomaxprocs=4 idleprocs=0 threads=16 spinningthreads=0 idlethreads=3 runqueue=29 [6 4 3 3]
SCHED 18096ms: gomaxprocs=4 idleprocs=2 threads=17 spinningthreads=1 idlethreads=5 runqueue=74 [1 0 0 0]
SCHED 19102ms: gomaxprocs=4 idleprocs=0 threads=19 spinningthreads=1 idlethreads=6 runqueue=3 [0 0 0 2]
SCHED 20103ms: gomaxprocs=4 idleprocs=0 threads=19 spinningthreads=1 idlethreads=6 runqueue=4 [0 0 0 0]
```

#### gc trace

```
 GODEBUG=gctrace=1 ./webapp > /dev/null
```

output

```
gc 99 @17.551s 3%: 1.8+4.8+0.054 ms clock, 7.2+0.32/1.8/0+0.21 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 100 @17.653s 3%: 18+12+0.056 ms clock, 73+1.0/1.5/0+0.22 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 101 @17.754s 3%: 1.2+2.4+0.075 ms clock, 5.0+1.1/0.84/0+0.30 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 102 @17.841s 3%: 2.3+2.3+0.078 ms clock, 9.2+0.61/1.1/0+0.31 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 103 @17.914s 3%: 14+9.4+0.057 ms clock, 59+0.70/1.0/0+0.22 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 104 @18.005s 3%: 1.2+8.5+0.053 ms clock, 4.9+8.9/1.4/0+0.21 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 105 @18.102s 3%: 29+10+0.061 ms clock, 118+0.31/1.5/0+0.24 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 106 @18.217s 3%: 4.6+37+0.077 ms clock, 18+0.27/6.7/0+0.30 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 107 @18.458s 3%: 0.53+1.4+0.19 ms clock, 2.1+2.1/0.87/0+0.76 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 108 @18.559s 3%: 43+5.6+0.082 ms clock, 174+2.0/1.7/0+0.33 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 109 @18.728s 3%: 0.17+2.4+0.45 ms clock, 0.68+0.28/1.4/0+1.8 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 110 @18.893s 3%: 0.043+1.2+0.065 ms clock, 0.17+1.2/0.77/0+0.26 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 111 @18.959s 3%: 0.15+0.88+0.062 ms clock, 0.60+1.3/0.63/0+0.25 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
gc 112 @19.055s 3%: 27+2.6+0.069 ms clock, 109+0.55/1.3/0+0.27 ms cpu, 4->4->1 MB, 5 MB goal, 4 P
GC forced
```

explain

gc #        the GC number, incremented at each GC
@#s         time in seconds since program start
`#%`          percentage of time spent in GC since program start
`#+...+#`     wall-clock/CPU times for the phases of the GC
`#->#->#` MB  heap size at GC start, at GC end, and live heap
`# MB goal`   goal heap size
`# P`         number of processors used

> The phases are stop-the-world (STW) sweep termination, concurrent
  mark and scan, and STW mark termination. The CPU times
  for mark/scan are broken down in to assist time (GC performed in
  line with allocation), background GC time, and idle GC time.
  If the line ends with "(forced)", this GC was forced by a
  runtime.GC() call and all phases are STW.

> you will see GC forced when cpu is idle to run, and gc will force itself to run

```
垃圾回收信息
gc 1 @2.104s 0%: 0.018+1.3+0.076 ms clock, 0.054+0.35/1.0/3.0+0.23 ms cpu, 4->4->3 MB, 5 MB goal, 4 P。
1 表示第一次执行
@2.104s 表示程序执行的总时间
0% 垃圾回收时间占用的百分比，（不知道和谁比？难道是和上面的程序执行总时间，这样比较感觉没意义）
0.018+1.3+0.076 ms clock 垃圾回收的时间，分别为0.018msSTW（stop-the-world）清扫的时间, 1.3ms并发标记和扫描的时间，0.076msSTW标记的时间
0.054+0.35/1.0/3.0+0.23 ms cpu 垃圾回收占用cpu时间, 0.054msSTW清扫占用的cpu时间，0.35ms辅助时间，1.0ms后台gc时间，3.0ms空闲gc时间，STW标记终止时间。 4->4->3 MB 堆的大小，gc后堆的大小，存活堆的大小
5 MB goal 整体堆的大小
4 P 使用的处理器数量
```

#### Memory Profiling

start pprof endpoint, see webapp/debug/debug.go

after start pprof endpoint, we can inspect pprof endpoint from `/debug/pprof`

to start an interactive terminal, 

```bash
go tool pprof -alloc_space http://localhost:8000/debug/pprof/allocs
```

```
(pprof) top 40 -cum
Showing nodes accounting for 268.58MB, 98.53% of 272.60MB total
Dropped 34 nodes (cum <= 1.36MB)
      flat  flat%   sum%        cum   cum%
         0     0%     0%   262.59MB 96.33%  net/http.(*conn).serve
   34.01MB 12.47% 12.47%   139.55MB 51.19%  net/http.(*conn).readRequest
         0     0% 12.47%   121.04MB 44.40%  net/http.HandlerFunc.ServeHTTP
         0     0% 12.47%   121.04MB 44.40%  net/http.serverHandler.ServeHTTP
    0.50MB  0.18% 12.66%   120.03MB 44.03%  github.com/xujihui1985/learninggo/profiling/webapp/service.tracing.func1.1
   21.51MB  7.89% 20.55%    78.52MB 28.80%  net/http.readRequest
   43.51MB 15.96% 36.51%    43.51MB 15.96%  net/textproto.(*Reader).ReadMIMEHeader
         0     0% 36.51%    40.01MB 14.68%  net/http.Header.Set
   40.01MB 14.68% 51.19%    40.01MB 14.68%  net/textproto.MIMEHeader.Set
```

```
(pprof) list nextRequestID
Total: 272.60MB
ROUTINE ======================== github.com/xujihui1985/learninggo/profiling/webapp/service.nextRequestID in /Users/sean/work/goworkspace/src/github.com/xujihui1985/learninggo/profiling/webapp/service/service.go
  512.01kB     2.50MB (flat, cum)  0.92% of Total
         .          .     60:           log.Fatalf("could not listen on %s, %v", host, err)
         .          .     61:   }
         .          .     62:}
         .          .     63:
         .          .     64:func nextRequestID() string {
  512.01kB     2.50MB     65:   return fmt.Sprintf("%d", time.Now().UnixNano())
         .          .     66:}
         .          .     67:
         .          .     68:func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
         .          .     69:   return func(next http.Handler) http.Handler {
         .          .     70:           return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
(pprof) web list nextRequestID
```

#### difference between allocs and heap

allocs default mode is allocs space, which is a sampling of all past memory allocations
heap default mode is inuse_spac, which is a sampling of memory allocations of live objects

#### cpu profile

```bash
go tool pprof http://localhost:8000/debug/pprof/profile\?seconds=5
```

#### execution tracing

```bash
go build
./tracing > t.out  // generate tracing file
go tool trace t.out // analize the tracing file
```



