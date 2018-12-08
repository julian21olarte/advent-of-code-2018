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

	inchMap := make([][]int, 1000)
	noOverlapingMap := make(map[string]*area)

	for i := 0; i < len(inchMap); i++ {
		inchMap[i] = make([]int, 1000)
	}

	// open file input
	file, err := os.Open("./../../inputs/day3.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function

	// generate the new scanner for read line by line the file
	scanner := bufio.NewScanner(file)

	// iterate in file opened
	for scanner.Scan() {
		line := scanner.Text() // read line
		var area area
		area, _ = getRectData(line)
		if err != nil {
			log.Fatal(err)
			return
		}
		applyOverlap(inchMap, &area, noOverlapingMap)
		
	}

	// iterate in the resulted matrix for count the overlaping
	inches, id := getOverlapingInchesAndNotOverlapingAreadID(inchMap, noOverlapingMap)
	fmt.Printf("Area not overlaping: " + id + "\n")
	fmt.Printf("Inches overlaping: " + inches)
}

// get overlaping inches and not overlaping area id
func getOverlapingInchesAndNotOverlapingAreadID(
	inchMap [][]int, 
	noOverlapingMap map[string] *area) (string, string){
		resp := 0	
		var id string
		var find = false
		for i := 0; i < len(inchMap); i++ {
			for j := 0; j < len(inchMap[i]); j++ {
				if inchMap[i][j] >= 2 {
					resp++
				}
				areaCheck, ok := noOverlapingMap[strconv.Itoa(i) + "-" + strconv.Itoa(j)]
				if !find && inchMap[i][j] == 1 && ok && !areaCheck.overlap {
					id = areaCheck.id
					find = true
				}
			}
		}
		return strconv.Itoa(resp), id
}


// applyOverlap function for apply the overlap in the specific submatrix
func applyOverlap(inchMap [][]int, area *area, noOverlapingMap map[string]*area) {
	var resp int
	for i := area.xDistance; i < (area.xDistance + area.width); i++ {
		for j := area.yDistance; j < (area.yDistance + area.height); j++ {
			inchMap[i][j]++
			if inchMap[i][j] >= 2 {
				area.overlap = true
				noOverlapingMap[strconv.Itoa(i) + "-" + strconv.Itoa(j)].overlap = true
				resp++
			} 
			if inchMap[i][j] == 1 {
				noOverlapingMap[strconv.Itoa(i) + "-" + strconv.Itoa(j)] = area
			}
		}	
	}
}

// getRectData function for transform the line string in a handler struct
func getRectData(line string) (area, error) {
	firstSplit := strings.Split(line, " @ ")
	key := firstSplit[1]
	id := firstSplit[0][1:]

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
		id: id,
		overlap: false,
	}, nil
}

// area struct for handle the coordenates and size of the rectangles
type area struct {
	xDistance int
	yDistance int
	width int
	height int
	key string
	id string
	overlap bool
}