package generator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	envNumTuples     = "SSIM_NUM_TUPLES"
	envAggSz         = "SSIM_AGGREGATE_SIZE"
	envFlushInterval = "SSIM_FLUSH_INTERVAL"

	defaultNumTuples     = 100
	defaultAggSz         = 2
	defaultFlushInterval = 0
)

func NumTuples() int {
	var numTuples int

	env := os.Getenv(envNumTuples)
	if len(env) == 0 {
		numTuples = defaultNumTuples
	} else {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envNumTuples)
		}

		if val <= 0 {
			log.Error.Fatalf("Invalid environment variable value: %s",
				envNumTuples)
		}
		numTuples = val
	}

	return numTuples
}

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

func FlushInterval() int {
	var flushInt int

	env := os.Getenv(envFlushInterval)
	if len(env) == 0 {
		flushInt = defaultFlushInterval
	} else {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envFlushInterval)
		}
		flushInt = val
	}

	return flushInt
}
