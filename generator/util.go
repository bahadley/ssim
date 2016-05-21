package generator

import (
	"math"
)

const (
	// Number of samples per second
	rate = 5000

	// Frequency
	hz = 261.625565
)

func GenSinData(samples int) []byte {
	f := (2.0 * math.Pi * hz) / rate
	data := make([]byte, samples, samples)
	for sample := 0; sample < samples; sample++ {
		data[sample] = byte(64.0 * (math.Sin(f*float64(sample)) + 1.0))
	}
	return data
}
