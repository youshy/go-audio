package sine

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func GenerateSine() {
	var (
		twopi float64
		steps int = 50
	)

	twopi = 2.0 * math.Pi // length of a single cycle of a sine wave
	angleincr := twopi / float64(steps)

	for i := 0; i < steps; i++ {
		samp := math.Sin(angleincr * float64(i))
		fmt.Printf("%.06f\n", samp)
	}
}

func GenerateTunedSine(freq float64) {
	var (
		twopi      float64
		samplerate int     = 44100
		duration   float64 = 2.0
		i          float64
	)

	steps := duration * float64(samplerate)
	twopi = 2.0 * math.Pi
	angleincr := twopi / float64(steps)

	for i = 0; i < steps; i++ {
		samp := math.Sin(angleincr * float64(i))
		fmt.Printf("%.06f\n", samp)
	}
}

func GenerateTunedSineToFile(freq float64, file string) error {
	var (
		duration   int     = 2
		samplerate int     = 44100
		twopi      float64 = math.Pi * 2.0

		start float64 = 1.0
		end   float64 = 1.0e-2
	)
	steps := duration * samplerate
	angleincr := twopi / float64(steps)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	decay := math.Pow(end/start, 1.0/float64(steps))
	for i := 0; i < steps; i++ {
		sample := math.Sin(angleincr * freq * float64(i))
		sample *= start
		start *= decay
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		_, err := f.Write(buf[:])
		if err != nil {
			return err
		}
	}
	return nil
}
