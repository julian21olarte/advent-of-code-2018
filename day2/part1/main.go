package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	var two, three int

	// open file input
	file, err := os.Open("./../../inputs/day2-12.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function

	// generate the new scanner for read line by line the file
	scanner := bufio.NewScanner(file)

	// iterate in file opened
	for scanner.Scan() {
		line := scanner.Text() // read line
		auxTwo, auxThree := getCount2and3repeats(line)

		if auxTwo > 0 {
			two += auxTwo
		}
		if auxThree > 0 {
			three += auxThree
		}
	}
	fmt.Printf(strconv.Itoa(two * three))
}

// get number of chars that repeat 2 and 3 times
func getCount2and3repeats(code string) (int, int) {
	var two, three int
	for _, char := range code {
		count := strings.Count(code, string(char))	
		if two < 1 && count == 2 {
			two++
		}
		if three < 1 && count == 3 {
			three++
		}
	}
	return two, three
}