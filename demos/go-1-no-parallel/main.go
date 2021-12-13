package main

import (
	"math"

	"github.com/dgravesa/go-openmp/demos/go-common"
)

func main() {
	common.PreLoop()

	// execute loop
	for i := 0; i < common.N; i++ {
		common.Results[i] = common.Sinc(common.Inputs[i] * math.Pi)
	}

	common.PostLoop()
}
