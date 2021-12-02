package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	"strings"
    "aoc2021/utils"
)

type position struct {
	x int
	y int
    aim int
}

type fn func(*position, int)

func forward(p *position, step int) {
	p.x += step
    p.y += p.aim * step
}

func down(p *position, step int) {
    p.aim += step
}

func up(p *position, step int) {
    p.aim -= step
}

func task2() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)

    moves := map[string] fn {
		"forward": forward,
		"down": down,
		"up": up,
	} 
	pos := position{0, 0, 0}
    
    for scanner.Scan() {
		line := strings.Fields(scanner.Text())
        step, _ := strconv.Atoi(line[1])

        moves[line[0]](&pos, step)
    }

    fmt.Println(pos.x * pos.y)
}

func main() {
    utils.Stopwatch(task2)
}