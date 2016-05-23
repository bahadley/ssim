package generator

const (
	sensor          = "A"
	measurementType = "T"
)

func Generate() ([]*SensorTuple, error) {
	qty := NumTuples()
	tuples := make([]*SensorTuple, qty)

	for i := 0; i < qty; i++ {
		tuples[i] = new(SensorTuple)
		tuples[i].Sensor = sensor
		tuples[i].Type = measurementType
	}

	GenSineData(tuples)

	return tuples, nil
}
