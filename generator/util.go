package generator

import (
	"math"
)

const (
	// Frequency
	hz = .2

	// Amplitude (peak-to-peak)
	amp = 128
)

func GenSineData(tuples []*SensorTuple) error {
	peakAmp := float64(amp / 2)

	for i := 0; i < len(tuples); i++ {
		tuples[i].Data = peakAmp * (math.Sin(hz*float64(i)) + 1.0)
		tuples[i].Data = roundDecimal(tuples[i].Data, 2)
	}

	return nil
}

func CalcAvg(tuples []*SensorTuple) error {
	aggSz := AggregateSize()
	sum := 0.0

	for i := 1; i <= len(tuples); i++ {
		sum += tuples[i-1].Data
		if i%aggSz == 0 {
			tuples[i-1].Aggregate = sum / float64(aggSz)
			tuples[i-1].Aggregate = roundDecimal(tuples[i-1].Aggregate, 2)
			sum = 0.0
		}
	}

	return nil
}

func roundDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
