package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"errors"
)

func main() {

	var codeArray []string

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
			codeArray = append(codeArray, line)
		}
		code, err := getCode(codeArray)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf(code)
}

// getCode function to find two strings with only one different character and return the string without the diff char
func getCode(codeArray []string) (string, error) {

	// search using the bubblesort theory O(n^2)
	for i := 0; i < len(codeArray); i++ {
		code := codeArray[i]
		for j := 1; j < len(codeArray); j++ {
			pos, count := countDiffChars(code, codeArray[j]) // count the diff chars inner the strings
			if count == 1 {
				return code[:pos] + code[pos + 1:], nil
			}
		}
	}
	return "", errors.New("Not exist match in codes")
}


// countDiffChars function to count the different chars inner 2 strings with the same length
func countDiffChars(code1 string, code2 string) (int, int) {
	var pos, count int
	for i := 0; i < len(code1); i++ {
		if code1[i] != code2[i] {
			count++
			pos = i
		}
	}
	if count == 1 {
		return pos, count
	}
	return 0, 0
}