package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"aoc2021/utils"
)

func CheckIncrease(prev int, curr int) int {
	if (curr - prev) > 0 {
		return 1
	}
	
	return 0
}

func task1() {
	file, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(file)

	var prev, curr int
	counter := 0

	scanner.Scan()
	prev, _ = strconv.Atoi(scanner.Text())
    
    for scanner.Scan() {
		curr, _ = strconv.Atoi(scanner.Text())

		counter += CheckIncrease(prev, curr)

		prev = curr
    }

	fmt.Println(counter)
}

func main() {
	utils.Stopwatch(task1)
}