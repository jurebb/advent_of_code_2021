package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "aoc2021/utils"
)

func task1() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var prev, curr int
    counter := 0

    scanner.Scan()
    prev, _ = strconv.Atoi(scanner.Text())
    
    for scanner.Scan() {
        curr, _ = strconv.Atoi(scanner.Text())

        if curr > prev {
            counter ++
        }

        prev = curr
    }

    fmt.Println(counter)
}

func main() {
    utils.Stopwatch(task1)
}