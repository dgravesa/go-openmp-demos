package main

import (
	"math"

	"github.com/dgravesa/go-openmp/demos/go-common"
	"github.com/dgravesa/go-parallel/parallel"
)

func main() {
	common.PreLoop()

	// execute loop
	parallel.For(common.N, func(i, _ int) {
		common.Results[i] = common.Sinc(common.Inputs[i] * math.Pi)
	})

	common.PostLoop()
}
