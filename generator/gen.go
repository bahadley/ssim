package generator

func Generate(qty int) ([]*SensorTuple, error) {
	tuples := make([]*SensorTuple, qty)

	for i := 0; i < qty; i++ {
		tuples[i] = new(SensorTuple)
		tuples[i].Sensor = "A"
	}

	return tuples, nil
}
