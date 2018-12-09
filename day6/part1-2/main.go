package main

import (
	"strconv"
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
	"time"
	"math"
	timeTracker "adventofcode/util"
)

func main() {

	defer timeTracker.TimeTrack(time.Now())
	file, err := os.Open("./../../inputs/day6.txt") // open file input
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function
	scanner := bufio.NewScanner(file)

	var coordArray []Coord
	var x, y int
	for scanner.Scan() {
		line := scanner.Text()
		newCoord := getCoord(&line)
		if newCoord.x > x {
			x = newCoord.x
		}
		if newCoord.y > y {
			y = newCoord.y
		}
		coordArray = append(coordArray, newCoord)
	}

	// create board with coordenates
	board := make([][]int, x + 1)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, y + 2)
	}

	fillBoard(&board, &coordArray)
	fmt.Printf("response 1: %d\n", getMaxArea(&board))
	//fmt.Print(infinites)
}

func getMaxArea(board *[][]int) int {
	infinites := getInfinites(board)
	maxMap := make(map[int]int)
	var max int
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			_, ok := infinites[(*board)[i][j]]
			if !ok && (*board)[i][j] != -1 {
				maxMap[(*board)[i][j]]++
			}
			if maxMap[(*board)[i][j]] > max {
				max = maxMap[(*board)[i][j]]
			}
		}	
	}
	return max
}
// fillBoard function to fill the board with the index of coords in area
func fillBoard(board *[][]int, coords *[]Coord) {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			(*board)[i][j] = markCoord(coords, i, j)
		}	
	}
}

// getInfinites function to get a map with the index of infinites coords
func getInfinites(board *[][]int) map[int]bool {
	mapInfinites := make(map[int]bool)
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if (i == 0 || j == 0 || i == len(*board) - 1 || j == len((*board)[i]) - 1) && (*board)[i][j] != -1 {
				mapInfinites[(*board)[i][j]] = true
			}
		}	
	}
	return mapInfinites
}

// markCoord function to mark in the board the specific coord
func markCoord(coords *[]Coord, x int, y int) int {
	aux := 10000
	auxIndex := 0
	var duplicate bool
	for index, coord := range *coords {
		auxX := int(math.Abs(float64(coord.x - x)))
		auxY := int(math.Abs(float64(coord.y - y)))
		sum := auxX + auxY
		if sum < aux {
			aux = sum
			auxIndex = index
			duplicate = false
		} else if sum == aux {
			duplicate = true
		}
	}
	if duplicate {
		return -1
	}
	return auxIndex
}

// getCoord func to create a new Coord struct
func getCoord(line *string) Coord {
	coords := strings.Split(*line, ", ")
	y, _ :=strconv.Atoi(coords[0])
	x, _ :=strconv.Atoi(coords[1])
	return Coord {
		x: x,
		y: y,
	}
}

// Coord struct
type Coord struct {
	x int
	y int
}