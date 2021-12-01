package utils

import (
	"fmt"
	"time"
)

func Stopwatch(function func()) {
	start := time.Now()

	function()

	elapsed := time.Since(start)
    fmt.Println("took ", elapsed)
}
