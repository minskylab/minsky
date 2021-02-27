package main

import (
	"fmt"
	"time"

	"github.com/minskylab/rca"
	"github.com/minskylab/rca/gol"
)

func main() {
	ca := gol.New(10, 10, time.Now().Unix())

	ticks := make(chan uint64)
	done := make(chan struct{})

	ca.RunInfiniteSimulation(ticks, done)

	go func(done chan struct{}) {
		time.Sleep(10 * time.Millisecond)
		done <- struct{}{}
	}(done)

	ca.Observe(ticks, func(n uint64, s rca.Space) {
		// c := s.(*gol.GoLCA)
		fmt.Println(n)
	})

}
