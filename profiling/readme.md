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

#### Block Profiling

