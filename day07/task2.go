package main

import (
    "aoc2021/utils"
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

const MAX_INT = int(^uint(0) >> 1)

func computeNeededFuel(optimalPos float64, candidateRange int, crabPositions []int) float64 {
    totalFuel := float64(MAX_INT)

    lowerBound := int(int(optimalPos) - (candidateRange))
    if lowerBound < 0 {
        lowerBound = 0
    }

    for i := lowerBound; i <= int(int(optimalPos)+(candidateRange)); i++ {

        totalFuelCandidate := 0.0
        for _, crabPos := range crabPositions {
            diff := math.Abs(float64(crabPos) - float64(i))
            totalFuelCandidate += diff * (diff + 1) / 2
        }

        if totalFuelCandidate < totalFuel {
            totalFuel = totalFuelCandidate
        }

    }

    return totalFuel
}

func computeOptimalPosition(crabPositions []int) (float64, int) {
    sort.Ints(crabPositions)

    len := len(crabPositions)
    maxPos := crabPositions[len-1]
    minPos := 0
    posRange := maxPos - minPos

    middlePos := float64(posRange) / 2.0
    optimalPos := 0.0

    for _, crabPos := range crabPositions {
        optimalPos += (middlePos - float64(crabPos))
    }

    optimalPos = math.Round(float64(optimalPos) / float64(posRange))

    return optimalPos, int(posRange / 2)
}

func initialize() []int {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var crabPositions []int

    scanner.Scan()
    initialFishes := strings.Split(scanner.Text(), ",")

    for _, fish := range initialFishes {
        fishNum, _ := strconv.Atoi(fish)
        crabPositions = append(crabPositions, fishNum)
    }

    return crabPositions
}

func task2() {
    initialCrabPositions := initialize()

    optimalPosCandidate, candidateRange := computeOptimalPosition(initialCrabPositions)
    fmt.Println("Optimal pos:", optimalPosCandidate)

    fuelNeeded := computeNeededFuel(optimalPosCandidate, candidateRange, initialCrabPositions)
    fmt.Println("Needed fuel:", int(fuelNeeded))
}

func main() {
    utils.Stopwatch(task2)
}
