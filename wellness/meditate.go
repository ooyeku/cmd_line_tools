package main

// command line tool for meditation and basic breathing exercises. By default, program will run 5 cycles of a 5 second inhale
// followed by a 5 second hold and a 5 second exhale.
// Flags: 	-n number of cycles
// 			-inhale number of seconds to inhale
// 			-exhale number of seconds to exhale
// 			-hold number of seconds to hold breathe between inhale and exhale
//
// Ex:
// go run meditate.go -n 3 -inhale 3 -exhale 5 -hold 4
// - will run a session with 3 rounds of 3 seconds inhales, a 4 second hold, then a 5 second exhale

import (
	"flag"
	"fmt"
	"time"
)

// getReady provides a count down to provide a buffer before the start of inhale/exhale session
func getReady() { // getting ready for breathing exercise
	fmt.Println("Relax and Prepare to inhale...")
	time.Sleep(time.Second * 5)

}

// breatheIn takes inhale, exhale, and hold second times as integers and executes the main meditation loop.
func breatheIn(inhale, exhale, hold int) {
	// buffer to get ready
	time.Sleep(time.Second * 3)
	fmt.Println("Breathe In for ", inhale, " seconds...")
	for i := 0; i <= inhale-1; i++ {
		fmt.Println("...")
		time.Sleep(time.Second * 1)
	}

	fmt.Println("Now Hold in for", hold, " seconds...")
	time.Sleep(time.Second * time.Duration(hold))
	fmt.Println("Good!")
	time.Sleep(time.Second * 1)
	fmt.Println("Now Exhale for", exhale, " seconds...")
	for i := 0; i <= exhale-1; i++ {
		fmt.Println("...")
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Well Done!!")
}

// TODO: delete
func waiting() {
	for {
		time.Sleep(time.Second)
		fmt.Println("...")
		time.Sleep(time.Second)
		fmt.Println("...")
		time.Sleep(time.Second)
	}

}

func main() {
	// use flag -n # where # is number of rounds to loop
	numRou := flag.Int("n", 5, "number of rounds [an integer]") // 5 rounds by default
	// use flag -inhale # where # is the number of seconds to inhale
	inhale := flag.Int("inhale", 5, "length of inhales in seconds [an integer]")
	// use flag -exhale # where # is the number of seconds to exhale
	exhale := flag.Int("exhale", 5, "length of exhale in seconds")
	// use flag -hold # where # is the number of seconds to hold breathe
	hold := flag.Int("hold", *inhale, "length to hold breathe [defaults to inhale time]")
	flag.Parse()
	for i := 1; i <= *numRou; i++ {
		go getReady()
		//	go waiting()
		breatheIn(*inhale, *exhale, *hold)
	}
}
