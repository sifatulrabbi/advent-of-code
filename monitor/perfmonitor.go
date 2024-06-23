package monitor

import (
	"fmt"
	"runtime"
	"time"
)

func PerfMonitor[A comparable, R comparable](fn func(A) R, args A) R {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
	start := time.Now()

	res := fn(args)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)

	return res
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
