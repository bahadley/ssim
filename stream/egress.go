package stream

import (
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/bahadley/ssim/generator"
	"github.com/bahadley/ssim/log"
)

var (
	send bool

	wg sync.WaitGroup
)

func Transmit(tuples []*generator.SensorTuple) {
	addrs := DstAddr()

	// Make a slice of channels and initialize them.
	chans := make([]chan *generator.SensorTuple, len(addrs))
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan *generator.SensorTuple, ChannelBufSz())
	}

	// Counting semaphore set to the number of channels.
	wg.Add(len(chans))

	// Launch all threads.  Each thread has a different destination.
	for idx, pipe := range chans {
		go egress(addrs[idx], pipe)
	}

	delayInt := DelayInterval()
	flushDelay := FlushDelay()

	// Send all the tuples by randomly selecting channels.
	for _, tuple := range tuples {
		if tuple.Type == generator.FlushType {
			time.Sleep(flushDelay * time.Millisecond)
		}
		chans[rand.Intn(len(chans))] <- tuple
		if tuple.Type == generator.FlushType {
			time.Sleep(flushDelay * time.Millisecond)
		} else {
			time.Sleep(delayInt * time.Millisecond)
		}
	}

	// Send a poison pill into the channels to shut down the threads.
	for _, pipe := range chans {
		pipe <- nil
	}

	// Wait for the threads to finish any queued work.
	wg.Wait()
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
	defer wg.Done()

	for {
		tuple := <-pipe
		if tuple == nil {
			// Poison pill found.
			return
		}

		msg, err := generator.Marshal(tuple)
		if err != nil {
			continue
		}

		log.Trace.Printf("Tx(%s): %s", dstAddr, msg)

		if !send {
			continue
		}

		_, err = conn.Write(msg)
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}

func init() {
	send = Send()
}
