package main

import (
	"fmt"
	"github.com/byt3smith/fir-go"
)

func main() {
	base := "https://firoku.herokuapp.com"
	token := "4dcbbbd7d865f2647f281945ce8465c8eda9f481"

	client := firGo.NewFIRClient(base, token)

	incidents, err := firGo.ListIncidents(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(incidents["results"], "\n")

	artifacts, err := firGo.ListArtifacts(client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(artifacts["results"], "\n")
}
