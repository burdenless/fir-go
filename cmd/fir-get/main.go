package main

import (
	"fmt"
	"os"

	"github.com/byt3smith/fir-go"
)

func main() {
	base := os.Getenv("FIR_BASE_URL")
	token := os.Getenv("FIR_APIKEY")

	// Instantiate new FIR client
	client := firGo.NewFIRClient(base, token)

	// Get incidents
	incidents, err := client.Incidents.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current Incident Listing:")
	fmt.Println(incidents, "\n")

	// Get artifacts
	artifacts, err := client.Artifacts.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Current Artifact Listing:")
	fmt.Println(artifacts, "\n")

	// Get users
	users, err := client.Users.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Current User Listing:")
	fmt.Println(users, "\n")
}
