package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	_ "strconv"
	"strings"
	_ "errors"
	"time"
	_	"github.com/golang-collections/collections/stack"
	timeTracker "adventofcode/util"
)

func main() {

	defer timeTracker.TimeTrack(time.Now())

	file, err := os.Open("./../../inputs/day5-12.txt") // open file input
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function
	scanner := bufio.NewScanner(file)

	// iterate in file opened	
	scanner.Scan()
	line := scanner.Text() // read line
	resp := replacePolymer(line)
	//fmt.Printf("%s\n", resp)
	fmt.Printf("%d\n", len(resp))
}

// replacePolymer function to search adjacent polymers that react and delete there ('c' with 'C' react, but 'C' with 'C' not react)
func replacePolymer(line string) string {
	var stack []string
	for i := 0; i < len(line); i++ {
		lenght := len(stack) - 1

		if len(stack) > 0 && reactPolymers(string(stack[lenght]), string(line[i])) {
			stack = stack[:lenght]
			for (len(stack) > 0) && (i < len(line) - 1) {
				i++
				pop := stack[len(stack) - 1]
				if reactPolymers(pop, string(line[i])) {
					stack = stack[:len(stack) - 1]
				} else {
					stack = append(stack, string(line[i]))
					break;
				}
			}
		} else {
			stack = append(stack, string(line[i]))
		}
	}
	return strings.Join(stack, "")
}

// reactPolymers function to check if 2 polymers react
func reactPolymers(a string, b string) bool {
	if a != b && strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}