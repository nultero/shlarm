package main

import (
	"time"
)

func main() {
	for {
		now := time.Now()
		printTime(now)
		time.Sleep(2 * time.Second) // tmp sleep
		// TODOOOO fix the face sliding w/ a lib like tcell
		// TODOOOOOO don't repaint every 2 seconds -- keep poll on minute channel?
	}
}

// char sauce
//https://en.wikipedia.org/wiki/Box-drawing_character
