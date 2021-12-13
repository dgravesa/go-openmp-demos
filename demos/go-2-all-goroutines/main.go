package main

import (
	"math"
	"sync"

	"github.com/dgravesa/go-openmp/demos/go-common"
)

func main() {
	common.PreLoop()

	// execute loop
	var wg sync.WaitGroup
	wg.Add(common.N)
	for i := 0; i < common.N; i++ {
		go func(i int) {
			defer wg.Done()
			common.Results[i] = common.Sinc(common.Inputs[i] * math.Pi)
		}(i)
	}
	wg.Wait()

	common.PostLoop()
}
