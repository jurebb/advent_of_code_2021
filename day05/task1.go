package main

import (
	"aoc2021/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type pointLine struct {
	a point
	b point
}

func generateAllPoints(pointLines []pointLine) (map[point]int, int) {
	pointcounter := map[point]int{}
	var lowerBound, upperBound int

	var countMoreThanTwo int

	for _, pl := range pointLines {
		if pl.a.x == pl.b.x {
			if pl.a.y > pl.b.y {
				lowerBound = pl.b.y
				upperBound = pl.a.y
			} else {
				lowerBound = pl.a.y
				upperBound = pl.b.y
			}

			for i := lowerBound; i <= upperBound; i++ {
				pointcounter[point{pl.a.x, i}] += 1

				if pointcounter[point{pl.a.x, i}] == 2 {
					countMoreThanTwo++
				}
			}
		} else if pl.a.y == pl.b.y {
			if pl.a.x > pl.b.x {
				lowerBound = pl.b.x
				upperBound = pl.a.x
			} else {
				lowerBound = pl.a.x
				upperBound = pl.b.x
			}

			for i := lowerBound; i <= upperBound; i++ {
				pointcounter[point{i, pl.a.y}] += 1

				if pointcounter[point{i, pl.a.y}] == 2 {
					countMoreThanTwo++
				}
			}
		}
	}

	return pointcounter, countMoreThanTwo
}

func task1() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pointLines []pointLine

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		x1, _ := strconv.Atoi(string(strings.Split(line[0], ",")[0]))
		y1, _ := strconv.Atoi(string(strings.Split(line[0], ",")[1]))
		x2, _ := strconv.Atoi(string(strings.Split(line[2], ",")[0]))
		y2, _ := strconv.Atoi(string(strings.Split(line[2], ",")[1]))

		if (x1 != x2) && (y1 != y2) {
			continue
		}

		point1 := point{x1, y1}
		point2 := point{x2, y2}

		pl := pointLine{point1, point2}
		pointLines = append(pointLines, pl)
	}

	_, count := generateAllPoints(pointLines)

	fmt.Println(count)
}

func main() {
	utils.Stopwatch(task1)
}
