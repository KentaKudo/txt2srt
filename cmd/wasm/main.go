package main

import (
	"bytes"
	"strings"
	"syscall/js"
	"time"

	"github.com/KentaKudo/txt2srt"
)

func main() {
	c := make(chan struct{})

	js.Global().Set("txt2srt", js.FuncOf(run))

	<-c
}

func run(_ js.Value, args []js.Value) any {
	in := args[0].String()
	r := strings.NewReader(in)

	var durationOption txt2srt.ConverterOption
	if len(args) == 2 {
		duration := args[1].Int()
		durationOption = txt2srt.WithSecondsToAdd(time.Duration(duration) * time.Second)
	}

	var buf bytes.Buffer
	txt2srt.Convert(r, &buf, durationOption)

	return buf.String()
}
