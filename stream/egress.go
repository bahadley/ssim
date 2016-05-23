package stream

import (
	"net"
	"time"

	"github.com/bahadley/ssim/generator"
	"github.com/bahadley/ssim/log"
	"github.com/bahadley/ssim/system"
)

var (
	transmit bool
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
		msg, err := generator.Marshal(tuple)
		if err != nil {
			continue
		}

		log.Trace.Printf("Tx(%s): %s", dstAddr, msg)

		if transmit {
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Warning.Println(err.Error())
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	transmit = system.Transmit()
}
