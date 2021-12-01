package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "aoc2021/utils"
)

const windowLen = 3

func CheckIncrease(prev int, curr int) int {
    if curr > prev {
        return 1
    }
    
    return 0
}

func task2() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()
    
    prevWindow := make([]int, 0)
    prevSum := int(0)
    currSum := int(0)
    windowCounter := 0
    counter := 0
    var curr int 

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        curr, _ = strconv.Atoi(scanner.Text())
        prevWindow = append(prevWindow, curr)
        prevSum += curr

        windowCounter += 1

        if windowCounter >= windowLen {
            currSum = prevSum
            break
        }
    }
    
    for scanner.Scan() {
        curr, _ = strconv.Atoi(scanner.Text())

        currSum -= prevWindow[0]
        prevWindow = prevWindow[1:]
        
        currSum += curr
        prevWindow = append(prevWindow, curr)

        counter += CheckIncrease(prevSum, currSum)

        prevSum = currSum
    }

    counter += CheckIncrease(prevSum, currSum)

    fmt.Println(counter)
}

func main() {
    utils.Stopwatch(task2)
}