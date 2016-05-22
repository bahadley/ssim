package stream

import (
	"fmt"
	"net"
	"time"

	"github.com/bahadley/ssim/generator"
	"github.com/bahadley/ssim/log"
	"github.com/bahadley/ssim/system"
)

const (
	msgFmt = "{\"sensor\":\"%s\",\"type\":\"%s\",\"ts\":%d,\"data\":%.2f}"
)

func Transmit(tuples []*generator.SensorTuple) {
	dstAddr, err := net.ResolveUDPAddr("udp",
		system.DstAddr()+":"+system.DstPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err := net.ResolveUDPAddr("udp",
		system.Addr()+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	for _, tuple := range tuples {
		msg := fmt.Sprintf(msgFmt, tuple.Sensor, "T", time.Now().UnixNano(), 40.4)

		log.Trace.Printf("Tx(%s): %s", dstAddr, msg)

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}
