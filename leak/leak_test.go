package main

import (
	"fmt"
	"testing"
)

func TestNativeLeak(t *testing.T) {
	for i := range 10 {
		fmt.Println("STEP:", i+1)
		processImage()
		printMemUsage()
	}
}