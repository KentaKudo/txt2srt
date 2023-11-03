package main

import (
	"bytes"
	"strings"
	"syscall/js"

	"github.com/KentaKudo/txt2srt"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("txt2srt", js.FuncOf(run))

	<-c
}

func run(_ js.Value, args []js.Value) any {
	in := args[0].String()
	r := strings.NewReader(in)

	var buf bytes.Buffer
	txt2srt.Convert(r, &buf)

	return buf.String()
}
