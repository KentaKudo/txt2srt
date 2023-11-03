package txt2srt

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

const format = "15:04:05,000"
const defaultSecondsToAdd = 3 * time.Second

func Convert(r io.Reader, w io.Writer, options ...ConverterOption) {
	c := converter{secondsToAdd: defaultSecondsToAdd}

	for _, opt := range options {
		opt(&c)
	}

	c.convert(r, w)
}

type converter struct {
	secondsToAdd time.Duration
}

type ConverterOption = func(c *converter)

func WithSecondsToAdd(secondsToAdd time.Duration) ConverterOption {
	return func(c *converter) {
		c.secondsToAdd = secondsToAdd
	}
}

func (c converter) convert(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var b strings.Builder

	idx := 1
	now, err := time.Parse(format, "00:00:00,000")
	if err != nil {
		log.Fatalf("time.Parse: %v", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		b.WriteString(fmt.Sprintf("%d\n", idx))
		b.WriteString(fmt.Sprintf("%v --> %v\n", now.Format(format), now.Add(c.secondsToAdd).Format(format)))
		b.WriteString(line)
		b.WriteString(fmt.Sprintln())
		b.WriteString(fmt.Sprintln())

		idx++
		now = now.Add(c.secondsToAdd)
	}

	fmt.Fprintln(w, b.String())
}
