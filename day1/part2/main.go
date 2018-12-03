package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
)

func main() {

	var resp, index int // initialize in 0 for default
	var auxArray []int
	intMap := make(map[int]int)

	// open file input
	file, err := os.Open("./../../inputs/day1-12.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // close file on finish main function

		// generate the new scanner for read line by line the file
    scanner := bufio.NewScanner(file)

		intMap[resp] = 1 // first current frecuency 0

		// iterate in file opened
    for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text()) // read line and convert string to int
			if err != nil {
				log.Fatal(err)
			}

			auxArray = append(auxArray, number) // storage the readed data for post iterate
			resp += number

			if checkMap(intMap, resp) {
				return
			}
		}

		// iterate the file data storaged
		for {
			resp += auxArray[index]
			index++
			if checkMap(intMap, resp) {
				return
			}
			if index == len(auxArray) {
				index = 0;
			}
		}
}

// checkMap function for check if the map contain the value passed
func checkMap(intMap map[int]int, value int) bool {
	_, exist := intMap[value]
	if exist {
		fmt.Printf(strconv.Itoa(value))				
		return true
	}
	intMap[value] = 1
	return false
}