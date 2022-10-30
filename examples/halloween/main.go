package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmowcomber/neopixel-client/client"
)

func main() {
	address := "http://192.168.4.1"
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	cli := client.New(address, httpClient)

	_, err := cli.SetBrightness(255)
	if err != nil {
		// log and ignore errors
		log.Printf("failed set brightness: %s", err.Error())
	}

	// set color to orange
	_, err = cli.SetColor("fc4605")
	if err != nil {
		log.Printf("failed set color: %s", err.Error())
	}

	interval := 5 * time.Second
	modeIndex := 0
	modes := []uint16{
		client.ModeBreath,
		client.ModeColorWipe,
		client.ModeHalloween,
	}
	for {
		mode := modes[modeIndex]
		log.Printf("Setting mode to %d", mode)
		_, err = cli.SetMode(mode)
		if err != nil {
			log.Printf("failed set mode: %s", err.Error())
		}
		time.Sleep(interval)

		modeIndex++
		if modeIndex > len(modes)-1 {
			modeIndex = 0
		}
	}
}
