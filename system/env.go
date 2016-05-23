package system

import (
	"os"
	"strings"
)

const (
	envTrace = "SSIM_TRACE"

	traceFlag = "YES"
)

func Trace() bool {
	t := os.Getenv(envTrace)
	if len(t) > 0 && strings.ToUpper(t) == traceFlag {
		return true
	} else {
		return false
	}
}
