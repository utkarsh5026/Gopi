package utils

import (
	"runtime"
	"time"
)

func BenchMark(fn func(), numThreads int) {
	start := time.Now()
	runtime.GOMAXPROCS(numThreads)
	fn()

	elapsed := time.Since(start)
	println("Execution time:", elapsed.Milliseconds(), "ms")
}
