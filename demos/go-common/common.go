package common

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var (
	N      int
	NumCPU int

	Inputs  []float64
	Results []float64

	seed        int64
	printSome   int
	runSerial   bool
	traceName   string
	cpuProfName string

	loopStartTime  time.Time
	loopFinishTime time.Time
)

func parseFlags() {
	flag.IntVar(&N, "N", 10000000, "number of work items")
	flag.IntVar(&NumCPU, "numgr", runtime.NumCPU(), "number of goroutines to use in parallel loop")
	flag.Int64Var(&seed, "seed", 1, "random generator seed")
	flag.IntVar(&printSome, "printsome", 0, "print first num values of result to verify")
	flag.BoolVar(&runSerial, "serial", false, "run non-parallelized for loop instead")
	flag.StringVar(&traceName, "trace", "", "output trace of loop to file")
	flag.StringVar(&cpuProfName, "cpuprofile", "", "output CPU profile of loop to file")
	flag.Parse()

	runtime.GOMAXPROCS(NumCPU)
}

func initArrays() {
	// initialize input array of N values
	rand.Seed(seed)
	Inputs = make([]float64, N)
	for i := 0; i < N; i++ {
		Inputs[i] = 10 * (rand.Float64() - 0.5) // -5 to 5
	}

	// allocate output array
	Results = make([]float64, N)
}

func printResults() {
	// print execution time
	totalLoopDuration := loopFinishTime.Sub(loopStartTime)
	fmt.Println(totalLoopDuration)

	// print some output values
	if printSome > 0 {
		fmt.Println("inputs:", Inputs[:printSome])
		fmt.Println("outputs:", Results[:printSome])
	}
}

func PreLoop() {
	parseFlags()

	initArrays()

	loopStartTime = time.Now()
}

func PostLoop() {
	loopFinishTime = time.Now()

	printResults()
}

func Sinc(x float64) float64 {
	if x == 0.0 {
		return 1.0
	}
	return math.Sin(x) / x
}
