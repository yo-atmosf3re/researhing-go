package lessons

import (
	"fmt"
	"researching-go/pkg/logger"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Printf(`miner came in mine number %d`, n)
	//time.Sleep(1 * time.Second)
	logger.Ptc("miner came out from mine")
	transferPoint <- 10
}

func mineCoal() {
	//transferPoint := make(chan int) // non-buffered channel, every goroutines wait when value will be taken from channel
	transferPoint := make(chan int, 3) // buffered channel. fixed quantity values which can put in this channel. goroutine wait when in channel will free place.
	var coal int

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	// if accesses channel more that completed goroutines that program is deadlock, because all goroutines is completed. in main goroutine main stream will be blocked, because value in channel is not exist. e.g.,if we have 3 goroutines and 4 access channel - this error is occurred.

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	time.Sleep(1200 * time.Millisecond)

	logger.Ptc("coal mined: ", coal)
	logger.Ptc("time has passed: ", time.Since(initTime))

}

func goroutineWithFor() {
	times := 1000000
	transferPoint := make(chan int, times)
	initTime := time.Now()
	coal := 0
	for i := 0; i < times; i++ {
		go mine(transferPoint, i)
		coal += <-transferPoint
	}
	logger.Ptc("coal mined: ", coal)
	logger.Ptc("time has passed: ", time.Since(initTime).Seconds())
}

func GcExample() {
	//mineCoal()
	//goroutineWithFor()
}
