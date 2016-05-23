package generator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	defaultAggSz = 2

	envAggSz = "SSIM_AGGREGATE_SIZE"
)

func AggregateSize() int {
	var aggSz int

	env := os.Getenv(envAggSz)
	if len(env) == 0 {
		aggSz = defaultAggSz
	} else {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envAggSz)
		}
		aggSz = val
	}

	return aggSz
}
