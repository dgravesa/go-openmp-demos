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
			// compute contiguous index range
			firstIndex, lastIndex := computeIndexBlock(grID, common.N, common.NumCPU)
			// execute loop on contiguous block
			for i := firstIndex; i < lastIndex; i++ {
				common.Results[i] = common.Sinc(common.Inputs[i] * math.Pi)
			}
		}(grID)
	}
	wg.Wait()

	common.PostLoop()
}

func computeIndexBlock(grID, N, numCPU int) (int, int) {
	div := N / numCPU
	mod := N % numCPU

	numWorkItems := div
	if grID < mod {
		numWorkItems++
	}

	firstIndex := grID*div + minInt(grID, mod)
	lastIndex := firstIndex + numWorkItems

	return firstIndex, lastIndex
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
