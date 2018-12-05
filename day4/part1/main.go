package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
	"strings"
	_ "errors"
	"time"
	"sort"
)
	
const timeFormat string = "2006-01-02 15:04"

func main() {

	var linesArray []*LineData
	var guardsArray []*GuardData
	guardsMap := make(map[string] *GuardData)
	file, err := os.Open("./../../inputs/day4-12.txt") // open file input
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // close file on finish main function
	scanner := bufio.NewScanner(file)

	// iterate in file opened	
	for scanner.Scan() {
		line := scanner.Text() // read line
		
		var newData, err = getLineData(line)
		linesArray = append(linesArray, &newData)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// sort the lines array by time
	sort.Slice(linesArray, func(i, j int) bool {
		return linesArray[i].time.Before(linesArray[j].time)
	})

	var currentGuardData *GuardData
	var guardMostSleep *GuardData // guard that most sleep
	guardMostSleepMinute := -1 // minute when the most sleep guard was sleep
	maxMinute := -1 // number of times that the guar was sleep in the mostSleepMinute
	sleepMinute := 0

	for _, line := range linesArray {
		_, min, _ := line.time.Clock()
		if line.text[0:1] == "G" {
			var newGuardData = getGuardData(line)
			existGuard, ok := guardsMap[newGuardData.id]
			if !ok {
				guardsArray = append(guardsArray, &newGuardData)
				currentGuardData = &newGuardData
				guardsMap[newGuardData.id] = &newGuardData
			} else {
				currentGuardData = existGuard
			}
		}
		if line.text[0:1] == "f" {
			sleepMinute = min
		}
		if line.text[0:1] == "w" {
				currentMinute := mapMinutesInGuardData(currentGuardData, sleepMinute, min)
				if currentGuardData.sleepMinutes > maxMinute {
					maxMinute = currentGuardData.sleepMinutes
					guardMostSleepMinute = currentMinute
					guardMostSleep = currentGuardData
				}
		}
	}
	idNumber, err := strconv.Atoi(guardMostSleep.id)
	if err != nil {
		log.Fatal("Error parsing the id")
	}
	fmt.Printf("GuardID: " + guardMostSleep.id + "\n")
	fmt.Printf("minute: " + strconv.Itoa(guardMostSleepMinute) + "\n")
	fmt.Printf("Response: " + strconv.Itoa(guardMostSleepMinute * idNumber))
}


func mapMinutesInGuardData(guardData *GuardData, sleep int, wake int) int {
	var max = -1
	var minute = -1
	for i := sleep; i < wake; i++ {
		guardData.minutesMap[i]++
		guardData.sleepMinutes++
		if guardData.minutesMap[i] > max {
			max = guardData.minutesMap[i]
			minute = i
		}
	}
	return minute
}

func getLineData(line string) (LineData, error) {
	timeString := line[1:17]
	time, err := time.Parse(timeFormat, timeString)
	if err != nil {
		return LineData{}, err
	}
	_, minute, _ := time.Clock()
	return LineData {
		time: time,
		text: line[19:],
		minute: minute,
	}, nil
}

func getGuardData(line *LineData) GuardData  {
	guardID := strings.Split(strings.Split(line.text, "#")[1], " ")[0]
	return GuardData {
		id: guardID,
		time: line.time,
		sleepMinutes: 0,
		minutesMap: make(map[int]int),
	}
}

// LineData struct for save the guard time and sleep data
type LineData struct {
	time time.Time	
	text string
	minute int
}
// GuardData struct for save the guard time and sleep data
type GuardData struct {
	id string
	time time.Time	
	sleepMinutes int
	minutesMap map[int]int
}