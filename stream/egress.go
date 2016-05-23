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
	addrs := system.DstAddr()

	for _, tuple := range tuples {
		Egress(addrs[0], tuple)
		time.Sleep(100 * time.Millisecond)
	}
}

func Egress(addr string, tuple *generator.SensorTuple) {
	dstAddr, err := net.ResolveUDPAddr("udp",
		addr+":"+system.DstPort())
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

	msg, err := generator.Marshal(tuple)
	//if err != nil {
	//	continue
	//}

	log.Trace.Printf("Tx(%s): %s", dstAddr, msg)

	if transmit {
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}

func init() {
	transmit = system.Transmit()
}
