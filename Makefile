.PHONY: run
run:
	go run cmd/web/main.go

txt2srt: cmd/txt2srt/main.go
	go build -o txt2srt cmd/txt2srt/main.go

txt2srt.wasm: cmd/wasm/main.go
	GOOS=js GOARCH=wasm go build -o docs/txt2srt.wasm cmd/wasm/main.go
