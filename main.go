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

	processData := mod.ExportedFunction("processData")
	printMemUsage := mod.ExportedFunction("printMemUsage")

	_, err = printMemUsage.Call(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mod.Memory().Size():", bToMb(mod.Memory().Size()), "MiB")

	for i := range 10 {
		fmt.Println("STEP:", i+1)
		step(ctx, mod, processData, printMemUsage)
	}
}

func step(ctx context.Context, mod api.Module, processData, printMemUsage api.Function) {
	_, err := processData.Call(ctx)
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
