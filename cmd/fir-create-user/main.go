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

	u := firGo.UserRequest{Username: "test"}

	err := client.Users.Create(&u)
	if err != nil {
		fmt.Println(err)
	}
}
