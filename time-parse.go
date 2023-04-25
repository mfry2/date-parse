package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	layout := time.RFC3339Nano
	timeNow := time.Now().UTC().Format(layout)
	timeOfTimeNow, err := time.Parse(layout, timeNow)
	timeTenAgo := time.Now().UTC().Add(time.Duration(-10) * time.Hour)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(timeOfTimeNow, timeTenAgo)
}
