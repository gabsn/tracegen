package main

import (
	"github.com/DataDog/dd-trace-go/tracer"
	"time"
)

func main() {
	for {
		span := tracer.NewRootSpan("test-go", "tracegen", "go")
		span.Finish()
		time.Sleep(1 * time.Second)
	}
}
