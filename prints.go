package main

import (
	"fmt"
	"shlarm/clock"
	"time"
)

func printTime(t time.Time) {

	// TODO the middle ':' in the clock face

	hour, minute := normalize(t)

	times := [][]string{clock.LeftBuf()} // start space to offset

	fmt.Printf("HOUR %v\n", hour)

	if hour < 10 {
		times = append(times, clock.NumsMap[0], clock.NumsMap[hour])
	} else {
		times = append(times, clock.NumsMap[1], clock.NumsMap[hour%10])
	}

	// minutes
	tensMin := minute / 10
	times = append(times, clock.NumsMap[tensMin])

	minute -= tensMin * 10
	times = append(times, clock.NumsMap[minute])

	strs := []string{}
	for i := 0; i < clock.Height; i++ {
		ln := ""
		for _, ti := range times {
			if len(ti) > 0 {
				ln += ti[i] + clock.Space
			}
		}
		strs = append(strs, ln)
	}

	fmt.Println("\033[2J")
	for _, s := range strs {
		fmt.Println(s)
	}
}

func normalize(t time.Time) (int, int) {
	if t.Hour() > 12 {
		return t.Hour() % 12, t.Minute()
	}

	return t.Hour(), t.Minute()
}
