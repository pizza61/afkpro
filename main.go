package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

func losu() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	c := rand.New(s1).Intn(60)
	return c
}

func main() {
	// flags
	tFlag := flag.Bool("t", false, "test mode (loops every 5 seconds instead of minutes)")
	flag.Parse()

	var dur time.Duration

	if *tFlag {
		fmt.Println(">> Test mode. Clicking every 5 seconds")
		dur = time.Second
	} else {
		fmt.Println(">> Clicking every 5-6 minutes")
		dur = time.Minute
	}

	// info
	fmt.Println(">> Ready. Your window has to be selected!")
	// ticker setup
	ticker := time.NewTicker(5 * dur)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				if !*tFlag {
					time.Sleep(time.Duration(losu()) * time.Second)
				}

				fmt.Println("Click")

				robotgo.KeyToggle("space", "down")
				time.Sleep(time.Second)
				robotgo.KeyToggle("space", "up")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// lock
	m := sync.Mutex{}
	m.Lock()
	m.Lock()
}
