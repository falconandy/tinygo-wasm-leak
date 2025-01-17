build-wasm:
	tinygo build -target=wasip1 -o ./leak/leak.wasm --no-debug -scheduler=none -buildmode=c-shared ./leak

run:
	go run .
