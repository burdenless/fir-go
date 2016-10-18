package main

import (
	"fmt"

	"github.com/byt3smith/fir-go"
)

func main() {

	client := firGo.NewFIRClient(nil, "https://firoku.herokuapp.com", "4dcbbbd7d865f2647f281945ce8465c8eda9f481")

	incidentlist, err := firGo.ListIncidents(client)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(incidentlist)
}
