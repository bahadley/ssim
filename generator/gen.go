package generator

const (
	sensor          = "A"
	temperatureType = "T"
	FlushType       = "F"
)

func Generate() ([]*SensorTuple, error) {
	qty := NumTuples()
	tuples := make([]*SensorTuple, qty)

	flushInt := FlushInterval()

	for i := 1; i <= qty; i++ {
		tuples[i-1] = new(SensorTuple)
		tuples[i-1].Sensor = sensor
		if flushInt > 0 && i%flushInt == 0 {
			tuples[i-1].Type = FlushType
		} else {
			tuples[i-1].Type = temperatureType
		}
	}

	GenSineData(tuples)

	return tuples, nil
}
