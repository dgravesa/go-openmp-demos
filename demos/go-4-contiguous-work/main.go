package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	var N int
	var seed int64
	var printSome int
	var numCPU int

	flag.IntVar(&N, "N", 1000000, "number of work items")
	flag.Int64Var(&seed, "Seed", 1, "random generator seed")
	flag.IntVar(&printSome, "PrintSome", 0, "print first num values of result to verify")
	flag.IntVar(&numCPU, "NumCPU", runtime.NumCPU(), "number of CPUs to use")
	flag.Parse()

	runtime.GOMAXPROCS(numCPU)

	// initialize input array of N values
	rand.Seed(seed)
	inputArray := make([]float64, N)
	for i := 0; i < N; i++ {
		inputArray[i] = 10 * (rand.Float64() - 0.5) // -5 to 5
	}

	// allocate output array
	outputArray := make([]float64, N)

	// execute loop
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(numCPU)
	for grID := 0; grID < numCPU; grID++ {
		go func(grID int) {
			defer wg.Done()
			// compute contiguous index range
			firstIndex, lastIndex := computeIndexBlock(grID, N, numCPU)
			// execute loop on contiguous block
			for i := firstIndex; i < lastIndex; i++ {
				outputArray[i] = sinc(inputArray[i] * math.Pi)
			}
		}(grID)
	}
	wg.Wait()
	stopTime := time.Now()

	// print execution time
	totalLoopDuration := stopTime.Sub(startTime)
	fmt.Println(totalLoopDuration)

	// print some output values
	if printSome > 0 {
		fmt.Println("inputs:", inputArray[:printSome])
		fmt.Println("outputs:", outputArray[:printSome])
	}
}

func sinc(x float64) float64 {
	if x == 0.0 {
		return 1.0
	}
	return math.Sin(x) / x
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
