package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
	"strings"
	"errors"
)

func main() {

	var resp int
	inchMap := make([][]int, 1050)
	searchMap := make(map[string]int)

	for i := 0; i < len(inchMap); i++ {
		inchMap[i] = make([]int, 1050)
	}

	// open file input
	file, err := os.Open("./../../inputs/day3-12.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function

	// generate the new scanner for read line by line the file
	scanner := bufio.NewScanner(file)

	// iterate in file opened
	for scanner.Scan() {
		line := scanner.Text() // read line
		area, err := getRectData(line)
		if err != nil {
			log.Fatal(err)
			return
		}
		applyOverlap(inchMap, area)
		searchMap[area.key] = searchMap[area.key] + 1
	}

	// iterate in the resulted matrix for count the overlaping
	resp = 0
	for i := 0; i < len(inchMap); i++ {
		for j := 0; j < len(inchMap[i]); j++ {
			if inchMap[i][j] >= 2 {
				resp++
			}
		}
	}
	fmt.Printf(strconv.Itoa(resp))
}


// applyOverlap function for apply the overlap in the specific submatrix
func applyOverlap(inchMap [][]int, area area) {
	var resp int
	for i := area.xDistance; i < (area.xDistance + area.width); i++ {
		for j := area.yDistance; j < (area.yDistance + area.height); j++ {
			inchMap[i][j]++
			if inchMap[i][j] >= 2 {
				resp++
			}
		}	
	}
}

// getRectData function for transform the line string in a handler struct
func getRectData(line string) (area, error) {
	key := strings.Split(line, "@ ")[1]
	auxLine := strings.Split(key, ",")
	xString, auxLine2 := auxLine[0], auxLine[1]
	xDistance, errx := strconv.Atoi(xString)

	auxLine = strings.Split(auxLine2, ": ")
	yString, auxLine4 := auxLine[0], auxLine[1]
	yDistance, erry := strconv.Atoi(yString)

	auxLine = strings.Split(auxLine4, "x")
	width, errw := strconv.Atoi(auxLine[0])
	height, errh := strconv.Atoi(auxLine[1])

	if errx != nil || erry != nil || errw != nil || errh != nil {
		return area{}, errors.New("Error parsing data to int")
	}
	return area {
		xDistance: xDistance,
		yDistance: yDistance,
		width: width,
		height: height,
		key: key,
	}, nil
}

// area struct for handle the coordenates and size of the rectangles
type area struct {
	xDistance int
	yDistance int
	width int
	height int
	key string
}