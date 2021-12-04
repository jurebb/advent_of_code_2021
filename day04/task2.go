package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "aoc2021/utils"
)

const boards = 100
const gridRowsCols = 5
var boardCount = 0

type number struct {
    x int
    y int
    board int
}

func parseActionsBoards(scanner *bufio.Scanner) ([]string, map[string] []number, [boards][]int) {
    scanner.Scan()
    actions := strings.Split(scanner.Text(), ",")

    rowCount := 0

    index := map[string] []number {}
    var boardValues [boards][]int
    
    for scanner.Scan() {
        if len(scanner.Text()) == 0 {
            continue
        }
        
        line := strings.Fields(scanner.Text())
        for pos, numStr := range line {
            index[numStr] = append(index[numStr], number{pos, rowCount, boardCount})
            numInt, _ := strconv.Atoi(numStr)
            boardValues[boardCount] = append(boardValues[boardCount], numInt)
        }

        rowCount++
        if rowCount >= gridRowsCols {
            rowCount=0
            boardCount++
        }
    }

    return actions, index, boardValues
}

func solve(actions []string, index map[string] []number, boardValues [boards][]int) {
    var drawnRows [boards][gridRowsCols][]int
    var drawnCols [boards][gridRowsCols][]int
    drawnBoardIndexesSet := make(map[string]bool)
    var solvedBoardsMap = map[string] bool{}

    for _, actChr := range actions {
        actInt, _ := strconv.Atoi(actChr)
        nums, ok := index[actChr]
        if ok {
            for _, num := range nums {
                drawnRows[num.board][num.x] = append(drawnRows[num.board][num.x], actInt)
                drawnCols[num.board][num.y] = append(drawnCols[num.board][num.y], actInt)
                drawnBoardIndexesSet[strconv.Itoa(num.board) + "," + strconv.Itoa(num.x) + "," + strconv.Itoa(num.y)] = true
            }
        }
        for i := 0; i < (boardCount - len(solvedBoardsMap)); i++ {
            solvedBoardIndex := checkSolved(drawnCols, solvedBoardsMap)
            if solvedBoardIndex > 0 {
                sum := sumUnmarked(solvedBoardIndex, drawnBoardIndexesSet, boardValues[solvedBoardIndex])
                if !solvedBoardsMap[strconv.Itoa(solvedBoardIndex)]{
                    fmt.Println("board ", solvedBoardIndex, " wins with drawn", actChr)
                    fmt.Println(sum * actInt)
                }
                
            }
            solvedBoardIndex = checkSolved(drawnRows, solvedBoardsMap)
            if solvedBoardIndex > 0 {
                sum := sumUnmarked(solvedBoardIndex, drawnBoardIndexesSet, boardValues[solvedBoardIndex])
                if !solvedBoardsMap[strconv.Itoa(solvedBoardIndex)]{
                    fmt.Println("board ", solvedBoardIndex, " wins with drawn", actChr)
                    fmt.Println(sum * actInt)
                }
                
            }

            if !solvedBoardsMap[strconv.Itoa(solvedBoardIndex)]{
                solvedBoardsMap[strconv.Itoa(solvedBoardIndex)] = true
            }
        }

        if len(solvedBoardsMap) >= boardCount -1 {
            break
        }
    }
}

func checkSolved(drawn [boards][gridRowsCols][]int, solvedBoardsMap map[string] bool) int {
    for boardIndex := range drawn {
        boardKey := strconv.Itoa(boardIndex)
        if solvedBoardsMap[boardKey] {
            continue
        }

        for rowColIndex := range drawn[boardIndex] {
            if len(drawn[boardIndex][rowColIndex]) == gridRowsCols {
                return boardIndex
            }
        }
    }
    
    return -1
}

func sumUnmarked(boardIndex int, drawnIndexesSet map[string]bool, boardValues []int) int {
    var unmarked []int
    boardKey := strconv.Itoa(boardIndex)
    
    for i := 0; i < gridRowsCols; i++ {
        for j:= 0; j < gridRowsCols; j++ {
            key := boardKey + "," + strconv.Itoa(i) + "," + strconv.Itoa(j)

            if drawnIndexesSet[key] {
                continue
            }

            index := j*gridRowsCols + i
            unmarked = append(unmarked, boardValues[index])
        }
    } 
    sum := 0
    for _, v := range unmarked {
      sum += v
    }
    return sum
}

func task2() {
    file, _ := os.Open(os.Args[1])
    defer file.Close()

    scanner := bufio.NewScanner(file)
    actions, index, boardValues := parseActionsBoards(scanner)

    solve(actions, index, boardValues)
}

func main() {
    utils.Stopwatch(task2)
}