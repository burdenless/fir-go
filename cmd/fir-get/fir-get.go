package main

import (
	"fmt"
	"github.com/byt3smith/fir-go"
)

func main() {
	base := "https://firoku.herokuapp.com"
	token := "4dcbbbd7d865f2647f281945ce8465c8eda9f481"

	// Instantiate new FIR client
	client := firGo.NewFIRClient(base, token)

	// Get incidents
	incidents, err := client.Incidents.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(incidents, "\n")

	// Get artifacts
	artifacts, err := client.Artifacts.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(artifacts, "\n")

	// Get users
	users, err := client.Users.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users, "\n")
}
