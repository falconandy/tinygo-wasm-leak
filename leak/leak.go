package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/png"
	"runtime"
)

//go:embed test-image.png
var testImageData []byte

//go:wasmexport processImage
func processImage() {
	img, err := png.Decode(bytes.NewReader(testImageData))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Decoded image: width=%d height=%d\n", img.Bounds().Dx(), img.Bounds().Dy())
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
