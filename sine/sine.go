package sine

import (
	"fmt"
	"math"
)

func generateSine(steps float64) {
	var (
		twopi float64
		i     float64
	)
	twopi = 2.0 * math.Pi
	angleincr := twopi / steps

	for i = 0; i < steps; i++ {
		samp := math.Sin(angleincr * i)
		fmt.Printf("%.02f\n", samp)
	}
}
