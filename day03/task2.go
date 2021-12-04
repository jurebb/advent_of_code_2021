package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "aoc2021/utils"
)

const lineLen = 12

func leastCommon(candidates []int, threshold int, totalLen int) ([]int, string) {
    var sliceZero []int 
    var sliceOne []int 
    countOne := 0

    for _, val := range candidates {
        if val < threshold {
            sliceZero = append(sliceZero, val)
        } else {
            sliceOne = append(sliceOne, val - threshold)
            countOne += 1
        }
    }

    if totalLen == 1 {
        if countOne > 0 {
            return sliceOne, "1"
        }
        return sliceZero, "0"
    }

    if countOne >= totalLen - countOne {
        return sliceZero, "0"
    }

    return sliceOne, "1"
}

func mostCommon(candidates []int, threshold int, totalLen int) ([]int, string) {
    var sliceZero []int 
    var sliceOne []int 
    countZero := 0

    for _, val := range candidates {
        if val >= threshold {
            sliceOne = append(sliceOne, val - threshold)
        } else {
            sliceZero = append(sliceZero, val)
            countZero += 1
        }
    }

    if totalLen == 1 {
        if countZero > 0 {
            return sliceZero, "0"
        }
        return sliceZero, "1"
    }

    if countZero > totalLen - countZero {
        return sliceZero, "0"
    }

    return sliceOne, "1"
}

func calculateCO2Rating(candidates []int, power int, binaryCO2Rating string) string {
    threshold := int(math.Pow(2, float64(power)))

    remainingLeastCommon, binaryVal := leastCommon(candidates, threshold, len(candidates))

    if power >= 0 {
        binaryCO2Rating = calculateCO2Rating(remainingLeastCommon, power - 1, binaryCO2Rating + binaryVal)
    }

    return binaryCO2Rating
}

func calculateO2Rating(candidates []int, power int, binaryO2Rating string) string {
    threshold := int(math.Pow(2, float64(power)))

    remainingMostCommon, binaryVal := mostCommon(candidates, threshold, len(candidates))

    if power >= 0 {
        binaryO2Rating = calculateO2Rating(remainingMostCommon, power - 1, binaryO2Rating + binaryVal)
    }

    return binaryO2Rating
}

func task2() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var candidates []int
    
    for scanner.Scan() {
        line := scanner.Text()
        decimal, _ := strconv.ParseInt(line, 2, 64)
        
        candidates = append(candidates, int(decimal))
    }

    binaryO2Rating := calculateO2Rating(candidates, lineLen - 1, "")
    binaryCO2Rating := calculateCO2Rating(candidates, lineLen - 1, "")

    o2Rating, _ := strconv.ParseInt(binaryO2Rating, 2, 64)
    co2Rating, _ := strconv.ParseInt(binaryCO2Rating, 2, 64)

    fmt.Println(o2Rating * co2Rating)
}

func main() {
    utils.Stopwatch(task2)
}