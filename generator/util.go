package generator

import (
	"math"
)

const (
	// Number of samples per second
	rate = 5000

	// Frequency
	hz = 260

	// Amplitude (peak-to-peak)
	amp = 128
)

func GenSineData(tuples []*SensorTuple) error {
	f := (2.0 * math.Pi * hz) / rate
	ub := float64(amp / 2)

	for i := 0; i < len(tuples); i++ {
		tuples[i].Data = ub * (math.Sin(f*float64(i)) + 1.0)
		tuples[i].Data = toFixed(tuples[i].Data, 2)
	}

	return nil
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
