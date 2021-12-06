package main

import (
	"aoc2021/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256 	// set to 80 for task 1

type FishBundle struct {
	count int
	countYoung int
}

func transferYoungFishes(day int, weekDayFish *map[int] FishBundle) {
	weekDay := day % 7
	newBirthday := (day + 7) % 7
	youngFishesCount := (*weekDayFish)[weekDay].countYoung

	(*weekDayFish)[weekDay] = FishBundle{ 
		(*weekDayFish)[weekDay].count, 
		0,
	}
	(*weekDayFish)[newBirthday] = FishBundle{ 
		(*weekDayFish)[newBirthday].count + youngFishesCount, 
		(*weekDayFish)[newBirthday].countYoung,
	}
}

func addUpcomingFishes(upcomingFishesCount int, day int, weekDayFish *map[int] FishBundle) {
	birthdayOfUpcoming := (day + 9) % 7
	(*weekDayFish)[birthdayOfUpcoming] = FishBundle{ 
		(*weekDayFish)[birthdayOfUpcoming].count, 
		(*weekDayFish)[birthdayOfUpcoming].countYoung + upcomingFishesCount,
	}
}

func simulateDays(numDays int, newFishes []int) {
	totalFishes := 0
	weekDayFish := map[int] FishBundle {}

	for _, fish := range newFishes {
		bundle := FishBundle{ weekDayFish[fish].count + 1, weekDayFish[fish].countYoung }
		weekDayFish[fish] = bundle

		totalFishes++
	}

	for day := 0; day < numDays; day++ {
		weekDay := day % 7
		
		upcomingFishesCount := weekDayFish[weekDay].count
		totalFishes += upcomingFishesCount
		
		addUpcomingFishes(upcomingFishesCount, day, &weekDayFish)

		transferYoungFishes(day, &weekDayFish)

		fmt.Println("End of day", day, ". Total fishes:", totalFishes)
	}
}

func initialize() []int {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var newbornFish []int

	scanner.Scan()
	initialFishes := strings.Split(scanner.Text(), ",")

	for _, fish := range initialFishes {
		fishNum, _ := strconv.Atoi(fish)
		newbornFish = append(newbornFish, fishNum)
	}

	return newbornFish
}

func task1() {	
	newFishes := initialize()

	simulateDays(DAYS, newFishes)
}

func main() {
	utils.Stopwatch(task1)
}
