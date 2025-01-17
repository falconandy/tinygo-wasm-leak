package main

import (
	_ "embed"
	"fmt"
	"runtime"
)

//go:wasmexport processData
func processData() {
	N := 1 * 1024 * 1024
	data := make([]byte, N)
	fmt.Println(len(data))
}

//go:wasmexport printMemUsage
func printMemUsage() {
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("runtime.MemStats: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB\n", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
