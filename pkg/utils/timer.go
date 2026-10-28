package utils

import (
	"log"
	"time"
)

func FunctionTimer(functionName string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		log.Printf("Took %s to execute function %s", elapsed, functionName)
	}
}
