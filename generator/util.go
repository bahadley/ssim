package generator

import (
	"math"
)

const (
	// Number of samples per second
	rate = 5000

	// Frequency
	hz = 261.625565

	// Amplitude
	amp = 128
)

func GenSinData(samples int) []byte {
	f := (2.0 * math.Pi * hz) / rate
	ub := float64(amp / 2)
	data := make([]byte, samples)
	for i := 0; i < samples; i++ {
		data[i] = byte(ub * (math.Sin(f*float64(i)) + 1.0))
	}
	return data
}
