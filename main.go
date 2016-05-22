package main

import (
	"os"
	"os/signal"
	"syscall"

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

	stream.Transmit()
}
