package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
	"time"
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
	resp1 := replacePolymer(line)
	_, resp2 := deleteMaxCountPolymer(&line)
	fmt.Printf("Response 1: %d\n", len(resp1))
	fmt.Printf("Response 2: %d\n", resp2)
}

func deleteMaxCountPolymer(line *string) (string, int) {
	polymer := ""
	countPolymer := 60000
	mapPolymers := make(map[string]int)
	for i := 0; i < len(*line); i++ {
		auxPolymer := strings.ToLower(string((*line)[i]))
		_, ok := mapPolymers[auxPolymer]
		if !ok {
			auxCount := getLenReplaceIgnoreCase(*line, auxPolymer)
			mapPolymers[auxPolymer] = auxCount
			if auxCount < countPolymer {
				countPolymer = auxCount
				polymer = auxPolymer
			}
		}
	}
	return polymer, countPolymer
}

// get the length of string after replaced ignored case polymer
func getLenReplaceIgnoreCase(line string, polymer string) int {
	line = strings.Replace(line, strings.ToLower(polymer), "", -1)
	line = strings.Replace(line, strings.ToUpper(polymer), "", -1)
	resp := replacePolymer(line)	
	return len(resp)
}

// replacePolymer function to search adjacent polymers that react and delete there ('c' with 'C' react, but 'C' with 'C' not react)
func replacePolymer(line string) (string) {
	var stack []string
	for i := 0; i < len(line); i++ {
		lenght := len(stack) - 1
		if len(stack) > 0 && checkIfReactPolymers(string(stack[lenght]), string(line[i])) {
			stack = stack[:lenght]
			for (len(stack) > 0) && (i < len(line) - 1) {
				i++
				pop := stack[len(stack) - 1]
				if checkIfReactPolymers(pop, string(line[i])) {
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

// checkIfReactPolymers function to check if 2 polymers react
func checkIfReactPolymers(a string, b string) bool {
	if a != b && strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}