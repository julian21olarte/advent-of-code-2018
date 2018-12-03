package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
)

func main() {

	var resp int

	// open file input
	file, err := os.Open("./../../inputs/day1-12.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // close file on finish main function

		// generate the new scanner for read line by line the file
    scanner := bufio.NewScanner(file)

		// iterate in file opened
    for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text()) // read line and convert string to int
			
			if err != nil {
				log.Fatal(err)
			}
			resp += number
		}
		fmt.Printf(strconv.Itoa(resp))
}