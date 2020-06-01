# golang channel 

写入 关闭的channel
> 测试表明当向关闭的channel中写入数据时会触发painc操作
> panic 类型: runtime.plainError
> panic 值: "send on closed channel"

读取 关闭的channel
> 当往关闭的channel中读数据时, 当读到第二个返回值为false时 就表示无数据输入了;
> v, ok := <- chan

unit test 
```
➜  go test -v -count=1 -test.failfast ./...
=== RUN   TestCloseWriteChanException
--- PASS: TestCloseWriteChanException (1.00s)
=== RUN   TestCloseReadChanException
--- PASS: TestCloseReadChanException (1.01s)
=== RUN   TestNoBufferChan
--- PASS: TestNoBufferChan (0.00s)
=== RUN   TestBufferChan
--- PASS: TestBufferChan (0.00s)
PASS
ok  	github.com/researchlab/gbp/channel	3.703s
?   	github.com/researchlab/gbp/channel/close	[no test files]
```

benchmark 
```
➜   go test -test.bench=".*" -count=5
goos: darwin
goarch: amd64
pkg: github.com/researchlab/gbp/channel
BenchmarkNoBufferChan-8     	  742285	      1541 ns/op
BenchmarkNoBufferChan-8     	  708918	      1510 ns/op
BenchmarkNoBufferChan-8     	  679736	      1537 ns/op
BenchmarkNoBufferChan-8     	  796756	      1507 ns/op
BenchmarkNoBufferChan-8     	  690397	      1524 ns/op
BenchmarkBufferChan100-8    	  641857	      1905 ns/op
BenchmarkBufferChan100-8    	  530427	      1907 ns/op
BenchmarkBufferChan100-8    	  560746	      1893 ns/op
BenchmarkBufferChan100-8    	  581718	      1922 ns/op
BenchmarkBufferChan100-8    	  632911	      1918 ns/op
BenchmarkBufferChan200-8    	  481228	      2330 ns/op
BenchmarkBufferChan200-8    	  480769	      2298 ns/op
BenchmarkBufferChan200-8    	  486703	      2309 ns/op
BenchmarkBufferChan200-8    	  463129	      2357 ns/op
BenchmarkBufferChan200-8    	  479805	      2295 ns/op
BenchmarkBufferChan500-8    	  334339	      3280 ns/op
BenchmarkBufferChan500-8    	  361581	      3255 ns/op
BenchmarkBufferChan500-8    	  336806	      3310 ns/op
BenchmarkBufferChan500-8    	  338631	      3255 ns/op
BenchmarkBufferChan500-8    	  324313	      3275 ns/op
BenchmarkBufferChan1000-8   	  205634	      4908 ns/op
BenchmarkBufferChan1000-8   	  233187	      4851 ns/op
BenchmarkBufferChan1000-8   	  235423	      4840 ns/op
BenchmarkBufferChan1000-8   	  249826	      4813 ns/op
BenchmarkBufferChan1000-8   	  238220	      4858 ns/op
PASS
ok  	github.com/researchlab/gbp/channel	46.266s
➜   go test -test.bench=".*" -cpuprofile=cpu.profile -test.memprofile=mem.profile --test.blockprofile=block.profile
goos: darwin
goarch: amd64
pkg: github.com/researchlab/gbp/channel
BenchmarkNoBufferChan-8     	  274166	      4462 ns/op
BenchmarkBufferChan100-8    	  230600	      4919 ns/op
BenchmarkBufferChan200-8    	  202717	      5275 ns/op
BenchmarkBufferChan500-8    	  178118	      6397 ns/op
BenchmarkBufferChan1000-8   	  142974	      8231 ns/op
PASS
ok  	github.com/researchlab/gbp/channel	10.366s
➜   go tool pprof channel.test cpu.profile
File: channel.test
Type: cpu
Time: Jun 1, 2020 at 10:36am (CST)
Duration: 9.22s, Total samples = 12.08s (131.00%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 11.49s, 95.12% of 12.08s total
Dropped 120 nodes (cum <= 0.06s)
Showing top 10 nodes out of 55
      flat  flat%   sum%        cum   cum%
     5.44s 45.03% 45.03%      5.44s 45.03%  runtime.pthread_cond_signal  # 唤醒阻塞线程
     3.09s 25.58% 70.61%      3.09s 25.58%  runtime.pthread_cond_wait # 用于阻塞当前线程
     1.68s 13.91% 84.52%      1.68s 13.91%  runtime.nanotime1  
     0.56s  4.64% 89.16%      0.56s  4.64%  runtime.usleep 
     0.26s  2.15% 91.31%      0.26s  2.15%  runtime.pthread_kill
     0.19s  1.57% 92.88%      0.27s  2.24%  runtime.scanobject  # 对象扫描 用于GC
     0.11s  0.91% 93.79%      0.11s  0.91%  runtime.(*waitq).dequeue
     0.07s  0.58% 94.37%      0.07s  0.58%  runtime.memmove
     0.07s  0.58% 94.95%      0.07s  0.58%  runtime.pthread_cond_timedwait_relative_np
     0.02s  0.17% 95.12%      0.67s  5.55%  runtime.gcDrain  #mark 标记
(pprof) exit
➜   go tool pprof channel.test mem.profile
File: channel.test
Type: alloc_space
Time: Jun 1, 2020 at 10:37am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 5135.68MB, 100% of 5138.13MB total
Dropped 21 nodes (cum <= 25.69MB)
      flat  flat%   sum%        cum   cum%
 2407.21MB 46.85% 46.85%  2407.21MB 46.85%  github.com/researchlab/gbp/channel.BenchmarkBufferChan1000
 1502.35MB 29.24% 76.09%  1502.35MB 29.24%  github.com/researchlab/gbp/channel.BenchmarkBufferChan500
  754.79MB 14.69% 90.78%   754.79MB 14.69%  github.com/researchlab/gbp/channel.BenchmarkBufferChan200
  379.32MB  7.38% 98.16%   379.32MB  7.38%  github.com/researchlab/gbp/channel.BenchmarkBufferChan100
   92.01MB  1.79%   100%    92.01MB  1.79%  github.com/researchlab/gbp/channel.BenchmarkNoBufferChan
         0     0%   100%  5135.68MB   100%  testing.(*B).launch
         0     0%   100%  5136.71MB   100%  testing.(*B).runN
(pprof) exit
➜   go tool pprof channel.test block.profile
File: channel.test
Type: delay
Time: Jun 1, 2020 at 10:37am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 13045.81ms, 100% of 13045.81ms total
Dropped 6 nodes (cum <= 65.23ms)
Showing top 10 nodes out of 33
      flat  flat%   sum%        cum   cum%
 9222.15ms 70.69% 70.69%  9222.15ms 70.69%  runtime.chanrecv1
 3823.66ms 29.31%   100%  3823.66ms 29.31%  runtime.chanrecv2
         0     0%   100%   428.82ms  3.29%  github.com/researchlab/gbp/channel.BenchmarkBufferChan100
         0     0%   100%   283.99ms  2.18%  github.com/researchlab/gbp/channel.BenchmarkBufferChan1000
         0     0%   100%   385.11ms  2.95%  github.com/researchlab/gbp/channel.BenchmarkBufferChan200
         0     0%   100%   335.31ms  2.57%  github.com/researchlab/gbp/channel.BenchmarkBufferChan500
         0     0%   100%   872.68ms  6.69%  github.com/researchlab/gbp/channel.BenchmarkNoBufferChan
         0     0%   100%  1005.17ms  7.70%  github.com/researchlab/gbp/channel.CloseReadChanException
         0     0%   100%   512.48ms  3.93%  github.com/researchlab/gbp/channel.Consumer
         0     0%   100%  1005.17ms  7.70%  github.com/researchlab/gbp/channel.TestCloseReadChanException
(pprof) exit
```

> pthread_cond_signal 函数的作用是发送一个信号给另外一个正在处于阻塞等待状态的线程,使其脱离阻塞状态,继续执行.如果没有线程处在阻塞等待状态,pthread_cond_signal也会成功返回。
>> 使用pthread_cond_signal一般不会有“惊群现象”产生，他最多只给一个线程发信号。假如有多个线程正在阻塞等待着这个条件变量的话，那么是根据各等待线程优先级的高低确定哪个线程接收到信号开始继续执行。如果各线程优先级相同，则根据等待时间的长短来确定哪个线程获得信号。但无论如何一个pthread_cond_signal调用最多发信一次。
>> 但是 pthread_cond_signal 在多处理器上可能同时唤醒多个线程，当你只能让一个线程处理某个任务时，其它被唤醒的线程就需要继续 wait，

> pthread_cond_wait()  用于阻塞当前线程，等待别的线程使用 pthread_cond_signal() 或pthread_cond_broadcast来唤醒它 。  

> usleep()函数是把调用该函数的线程挂起一段时间，单位是微秒（百万分之一秒）

go test  参数说明
```
➜   go test -b
flag provided but not defined: -b
Usage of /var/folders/5q/7gy5r2h91_s477x2g8vn30hc0000gn/T/go-build174561908/b001/channel.test:
  -test.bench regexp
    	run only benchmarks matching regexp
  -test.benchmem
    	print memory allocations for benchmarks
  -test.benchtime d
    	run each benchmark for duration d (default 1s)
  -test.blockprofile file
    	write a goroutine blocking profile to file
  -test.blockprofilerate rate
    	set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)
  -test.count n
    	run tests and benchmarks n times (default 1)
  -test.coverprofile file
    	write a coverage profile to file
  -test.cpu list
    	comma-separated list of cpu counts to run each test with
  -test.cpuprofile file
    	write a cpu profile to file
  -test.failfast
    	do not start new tests after the first test failure
  -test.list regexp
    	list tests, examples, and benchmarks matching regexp then exit
  -test.memprofile file
    	write an allocation profile to file
  -test.memprofilerate rate
    	set memory allocation profiling rate (see runtime.MemProfileRate)
  -test.mutexprofile string
    	write a mutex contention profile to the named file after execution
  -test.mutexprofilefraction int
    	if >= 0, calls runtime.SetMutexProfileFraction() (default 1)
  -test.outputdir dir
    	write profiles to dir
  -test.parallel n
    	run at most n tests in parallel (default 8)
  -test.run regexp
    	run only tests and examples matching regexp
  -test.short
    	run smaller test suite to save time
  -test.testlogfile file
    	write test action log to file (for use only by cmd/go)
  -test.timeout d
    	panic test binary after duration d (default 0, timeout disabled)
  -test.trace file
    	write an execution trace to file
  -test.v
    	verbose: print additional output

格式形如：
go test [-c] [-i] [build flags] [packages] [flags for test binary]

参数解读：
-c : 编译go test成为可执行的二进制文件，但是不运行测试。

-i : 安装测试包依赖的package，但是不运行测试。

关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空

关于packages，调用go help packages，这些是关于包的管理，一般设置为空

关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数

-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。

-test.run pattern: 只跑哪些单元测试用例

-test.bench patten: 只跑那些性能测试用例

-test.benchmem : 是否在性能测试的时候输出内存情况

-test.benchtime t : 性能测试运行的时间，默认是1s

-test.cpuprofile cpu.out : 是否输出cpu性能分析文件

-test.memprofile mem.out : 是否输出内存性能分析文件

-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件

-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。

你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。

-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下

-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。

-test.timeout t : 如果测试用例运行时间超过t，则抛出panic

-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理

-test.short : 将那些运行时间较长的测试用例运行时间缩短
```
