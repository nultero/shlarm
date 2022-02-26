package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playSound() {

	// obviously TODO to play sounds at a configurable time / amount

	f, err := os.Open("./testsound.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	prog := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		prog <- true
	})))

	<-prog
}
