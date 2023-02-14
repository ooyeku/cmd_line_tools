package main

import (
	"flag"
	"fmt"
	"time"
)

func animate(s string, duration int) {
	for i := 0; i < duration; i++ {
		fmt.Print(s)
		time.Sleep(time.Second)
	}
}

func ready() {
	fmt.Println("Relax and prepare to inhale...")
	time.Sleep(time.Second * 5)
}

func in(inhale, exhale, hold int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Breathe in for", inhale, "seconds...")
	animate(".", inhale)
	fmt.Println("\nNow hold for", hold, "seconds...")
	time.Sleep(time.Second * time.Duration(hold))
	fmt.Println("Good!")
	fmt.Println("Now exhale for", exhale, "seconds...")
	animate(".", exhale)
	fmt.Println("\nWell done!!")
}

func main() {
	numRounds := flag.Int("n", 5, "number of rounds")
	inhaleDuration := flag.Int("inhale", 5, "inhale duration in seconds")
	exhaleDuration := flag.Int("exhale", 5, "exhale duration in seconds")
	holdDuration := flag.Int("hold", *inhaleDuration, "hold duration in seconds")
	flag.Parse()

	for i := 1; i <= *numRounds; i++ {
		go ready()
		in(*inhaleDuration, *exhaleDuration, *holdDuration)
	}
}
