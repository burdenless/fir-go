package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/byt3smith/fir-go"
)

func main() {
	base := os.Getenv("FIR_BASE_URL")
	token := os.Getenv("FIR_APIKEY")

	// Instantiate new FIR client
	client := firGo.NewFIRClient(base, token)

	eventName := flag.String("event", "new-event", "Name of new event")
	flag.Parse()

	if *eventName != "new-event" {
		fmt.Println("Creating Event:", *eventName)
		i := firGo.IncidentRequest{Subject: *eventName, Description: *eventName}
		err := client.Incidents.Create(&i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
