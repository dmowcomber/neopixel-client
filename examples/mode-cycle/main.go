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

	respBody, err := cli.SetBrightness(110)
	if err != nil {
		log.Printf("failed set brightness: %s", err.Error())
	}
	log.Printf("response: %s", respBody)

	// cycle through all the modes
	var i uint16 = 0
	for ; i <= 64; i++ {
		respBody, err := cli.SetMode(i)
		if err != nil {
			log.Printf("failed to do http: %s", err.Error())
		}
		log.Printf("response: %s", respBody)
		time.Sleep(2 * time.Second)
	}
}
