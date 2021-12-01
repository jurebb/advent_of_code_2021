package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const window_len = 3

func sum(array []int64) int64 {  
	result := int64(0)

	for _, v := range array {  
		result += v  
	} 

	return result  
}  

func main() {
	file, _ := os.Open(os.Args[1])
	
	prev_window := make([]int64, 0)
	window_counter := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		prev_window = append(prev_window, curr)

		window_counter += 1

		if window_counter >= window_len {
			break
		}
	}
    
    for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 0, 64)

		prev_window = prev_window[1:]
		prev_window = append(prev_window, curr)

		fmt.Println(sum(prev_window))
    }

	// fmt.Println(counter)
}