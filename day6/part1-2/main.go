package main

import (
	_ "fmt"
	"log"
	"bufio"
	"os"
	_ "strings"
	"time"
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

	// iterate in file opened	
	scanner.Scan()
	line := scanner.Text() // read line

}
