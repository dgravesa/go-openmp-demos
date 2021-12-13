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
	wg.Add(common.NumCPU)
	for grID := 0; grID < common.NumCPU; grID++ {
		go func(grID int) {
			defer wg.Done()
			for i := grID; i < common.N; i += common.NumCPU {
				common.Results[i] = common.Sinc(common.Inputs[i] * math.Pi)
			}
		}(grID)
	}
	wg.Wait()

	common.PostLoop()
}
