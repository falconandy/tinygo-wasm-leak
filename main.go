package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed leak/leak.wasm
var leakWasm []byte

func main() {
	ctx := context.Background()

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.InstantiateWithConfig(ctx, leakWasm, wazero.NewModuleConfig().WithStartFunctions("_initialize").WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}

	processImage := mod.ExportedFunction("processImage")
	printMemUsage := mod.ExportedFunction("printMemUsage")

	for i := range 10 {
		fmt.Println("STEP:", i+1)
		step(ctx, mod, processImage, printMemUsage)
	}
}

func step(ctx context.Context, mod api.Module, processImage, printMemUsage api.Function) {
	_, err := processImage.Call(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = printMemUsage.Call(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mod.Memory().Size():", bToMb(mod.Memory().Size()), "MiB")

}

func bToMb(b uint32) uint32 {
	return b / 1024 / 1024
}
