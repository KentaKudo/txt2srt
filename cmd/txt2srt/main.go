package main

import (
	"os"

	"github.com/KentaKudo/txt2srt"
)

func main() {
	txt2srt.Convert(os.Stdin, os.Stdout)
}
