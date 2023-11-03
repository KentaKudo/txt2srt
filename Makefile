.PHONY: run
run:
	@go run main.go

txt2srt: cmd/txt2srt/main.go
	go build -o txt2srt cmd/txt2srt/main.go
