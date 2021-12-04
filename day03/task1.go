package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "aoc2021/utils"
)

const lineLen = 12

type fn func(*int) 

var moves = map[rune] fn {
    '0': func(x *int) { *x-- },
    '1': func(x *int) { *x++ },
}

func updateBitPosition(counter *int, bit rune) {
    moves[bit](counter)
}

func calculateGamma(commons [lineLen]int) int {
    gamma := 0
    for pos, val := range commons {
        if val > 0 {
            gamma += int(math.Pow(2, float64(lineLen - pos - 1)))
        }
    }

    return gamma
}

func calculateEpsilon(commons [lineLen]int) int {
    epsilon := 0
    for pos, val := range commons {
        if val < 0 {
            epsilon += int(math.Pow(2, float64(lineLen - pos - 1)))
        }
    }

    return epsilon
}

func task1() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)

    commons := [lineLen]int{}
    
    for scanner.Scan() {
        line := scanner.Text()
        for pos, char := range line {
            updateBitPosition(&commons[pos], char)
        }
    }

    gamma := calculateGamma(commons)
    epsilon := calculateEpsilon(commons)

    fmt.Println(gamma * epsilon)
}

func main() {
    utils.Stopwatch(task1)
}