package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bahadley/ssim/generator"
	"github.com/bahadley/ssim/log"
	"github.com/bahadley/ssim/stream"
)

func main() {
	log.Info.Println("Starting up ...")

	// Allow the node to be shut down gracefully.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block waiting for signal.
		<-c
		log.Info.Println("Shutting down ...")
		os.Exit(0)
	}()

	tuples, err := generator.Generate(200)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	stream.Transmit(tuples)
	generator.CalcAvg(tuples)

	for _, tuple := range tuples {
		fmt.Printf("%d,%.2f,%.2f\n", tuple.Timestamp, tuple.Data, tuple.Aggregate)
	}

	log.Info.Println("Shutting down ...")
}
