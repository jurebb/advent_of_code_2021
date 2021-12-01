package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(file)

	var prev = int64(-1)
	counter := 0
    
    for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 0, 64)

		if prev != -1 {
			if (curr - prev) > 0 {
				counter += 1
			}
		}

		prev = curr
    }

	fmt.Println(counter)
}