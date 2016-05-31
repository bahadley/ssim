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

	for i := 0; i < qty; i++ {
		tuples[i] = new(SensorTuple)
		tuples[i].Sensor = sensor
		if flushInt > 0 && i > 0 && i%flushInt == 0 {
			tuples[i].Type = FlushType
		} else {
			tuples[i].Type = temperatureType
		}
	}

	GenSineData(tuples)

	return tuples, nil
}
