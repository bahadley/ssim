package stream

import (
	"math/rand"
	"net"
	"time"

	"github.com/bahadley/ssim/generator"
	"github.com/bahadley/ssim/log"
)

var (
	send bool
)

func Transmit(tuples []*generator.SensorTuple) {
	addrs := DstAddr()

	chans := make([]chan *generator.SensorTuple, len(addrs))
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan *generator.SensorTuple, ChannelBufSz())
	}

	for idx, pipe := range chans {
		go egress(addrs[idx], pipe)
	}

	delayInt := DelayInterval()

	for _, tuple := range tuples {
		chans[rand.Intn(len(chans))] <- tuple
		time.Sleep(delayInt * time.Millisecond)
	}

	for _, pipe := range chans {
		close(pipe)
	}
}

func egress(addr string, pipe chan *generator.SensorTuple) {
	dstAddr, err := net.ResolveUDPAddr("udp",
		addr+":"+DstPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err := net.ResolveUDPAddr("udp",
		Addr()+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	for {
		tuple, ok := <-pipe
		if !ok {
			return
		}

		msg, err := generator.Marshal(tuple)
		if err != nil {
			continue
		}

		log.Trace.Printf("Tx(%s): %s", dstAddr, msg)

		if send {
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Warning.Println(err.Error())
			}
		}
	}
}

func init() {
	send = Send()
}
