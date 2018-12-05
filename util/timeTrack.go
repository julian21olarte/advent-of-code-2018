package util

import (
	"time"
	"log"
)

// TimeTrack function to track the execution time
func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Time Exec:  %s", elapsed)
}