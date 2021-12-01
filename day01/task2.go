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

func check_increase(prev int64, curr int64) int {
	if (curr - prev) > 0 {
		return 1
	}
	
	return 0
}

func main() {
	file, _ := os.Open(os.Args[1])
	
	prev_window := make([]int64, 0)
	prev_sum := int64(0)
	curr_sum := int64(0)
	window_counter := 0
	counter := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		prev_window = append(prev_window, curr)

		window_counter += 1

		if window_counter >= window_len {
			prev_sum = sum(prev_window)
			curr_sum = prev_sum
			break
		}
	}
    
    for scanner.Scan() {
		curr, _ := strconv.ParseInt(scanner.Text(), 0, 64)

		curr_sum -= prev_window[0]
		prev_window = prev_window[1:]
		
		curr_sum += curr
		prev_window = append(prev_window, curr)

		counter += check_increase(prev_sum, curr_sum)

		prev_sum = curr_sum
    }

	counter += check_increase(prev_sum, curr_sum)

	fmt.Println(counter)
}