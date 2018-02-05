package main

// https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html
import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	wg.Wait()

	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}

// go build example.go

// 处理器设置为1
// GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./example

// 结果
// SCHED 0ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 [9]

// SCHED 1000ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=1 runqueue=0 [9]

// 1000ms:程序启动后的时间。表示一秒钟
// gomaxprocs=1:配置的处理器数量
// idleprocs=0:不繁忙的处理器数量
// threads=3:运行时正在管理的线程数
// spinningthreads=0:
// idlethreads=1:不繁忙的线程数
// runqueue=0:全局运行队列中的goroutine数目
// [9]:本地运行队列中的goroutine数目

// 处理器设置为2
// GOMAXPROCS=2 GODEBUG=schedtrace=1000 ./example

// SCHED 2002ms: gomaxprocs=2 idleprocs=0 threads=4 spinningthreads=0 idlethreads=1 runqueue=0 [4 4]

// 2002ms        : This is the trace for the 2 second mark.
// gomaxprocs=2  : 2 processors are configured for this program.
// threads=4     : 4 threads exist. 2 for processors and 2 for the runtime.
// idlethreads=1 : 1 idle thread (3 threads running).
// idleprocs=0   : 0 processors are idle (2 processors busy).
// runqueue=0    : All runnable goroutines have been moved to a local run queue.
// [4 4]         : 4 goroutines are waiting inside each local run queue.

// 更详细的信息
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 ./example

// G运行时的状态
// status: http://golang.org/src/runtime/
// Gidle,            // 0
// Grunnable,        // 1 runnable and on a run queue
// Grunning,         // 2 running
// Gsyscall,         // 3 performing a syscall
// Gwaiting,         // 4 waiting for the runtime
// Gmoribund_unused, // 5 currently unused, but hardcoded in gdb scripts
// Gdead,            // 6 goroutine is dead
// Genqueue,         // 7 only the Gscanenqueue is used
// Gcopystack,       // 8 in this state when newstack is moving the stack
