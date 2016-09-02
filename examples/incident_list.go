package main

import (
	"fmt"

	"github.com/byt3smith/fir_go_client"
)

func main() {

	client := firGo.NewFIRClient(nil, "https://firoku.herokuapp.com")
	incidentlist, _, err := client.Incident.Get()
	if err != nil {
		fmt.Println(incidentlist)
	} else {
		fmt.Println(err)
	}
}
