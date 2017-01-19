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
	incidents, err := firGo.ListIncidents(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(incidents["results"], "\n")

	// Get artifacts
	artifacts, err := firGo.ListArtifacts(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(artifacts["results"], "\n")


	// Get users
	users, err := firGo.ListUsers(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users["results"], "\n")
}
