package system

import (
	"os"
	"strings"
)

const (
	envAddr     = "SSIM_ADDR"
	envDstAddr  = "SSIM_DST_ADDR"
	envDstPort  = "SSIM_DST_PORT"
	envTransmit = "SSIM_TRANSMIT"
	envTrace    = "SSIM_TRACE"

	defaultAddr    = "localhost"
	defaultDstAddr = "localhost"
	defaultDstPort = "22221"
	traceFlag      = "YES"
	noTransmitFlag = "NO"
)

func Addr() string {
	addr := os.Getenv(envAddr)
	if len(addr) == 0 {
		return defaultAddr
	} else {
		return addr
	}
}

func DstAddr() []string {
	addr := os.Getenv(envDstAddr)
	if len(addr) == 0 {
		return []string{defaultDstAddr}
	} else {
		return strings.Split(addr, ",")
	}
}

func DstPort() string {
	port := os.Getenv(envDstPort)
	if len(port) == 0 {
		return defaultDstPort
	} else {
		return port
	}
}

func Trace() bool {
	t := os.Getenv(envTrace)
	if len(t) > 0 && strings.ToUpper(t) == traceFlag {
		return true
	} else {
		return false
	}
}

func Transmit() bool {
	t := os.Getenv(envTransmit)
	if len(t) > 0 && strings.ToUpper(t) == noTransmitFlag {
		return false
	} else {
		return true
	}
}
