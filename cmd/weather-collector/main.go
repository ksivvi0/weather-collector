package main

import (
	"os"
	"os/signal"
	"syscall"
	"weather-collector/internal/collector"
)

func main() {
	c, err := collector.NewCollector()
	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	c.Run(sigChan)
}
