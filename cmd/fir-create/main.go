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

	u := firGo.User{
		Username: "joe",
	}

	err := client.Users.Create(&u)
	if err != nil {
		fmt.Println(err)
	}
}
