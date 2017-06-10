package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const usersPath = "/users"

// UsersInterface is an interface for all user struct methods
type UsersInterface interface {
	List() ([]User, error)
	Create(*UserRequest) error
}

// User is a model for FIR users
type User struct {
	ID       int    `json:"id",omitempty`
	Groups   []int  `json:"groups",omitempty`
	Email    string `json:"email",omitempty`
	Username string `json:"username",omitempty`
	URL      string `json:"url",omitempty`
}

// UserRequest type is User without an ID
type UserRequest struct {
	Email    string `json:"email",omitempty`
	Username string `json:"username",omitempty`
	URL      string `json:"url",omitempty`
}

// UserResponse holds a response from FIR
type UserResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []User
}

// UserServiceObj registers User as a service to the client
type UserServiceObj struct {
	client *Client
}

// List current FIR incidents
func (us *UserServiceObj) List() ([]User, error) {
	req, err := us.client.NewRequest("GET", usersPath, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := us.client.Do(req)

	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("ERROR.1 :", err2)
		}

		var dat UserResponse
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}

		return dat.Results, nil
	}

	fmt.Println("ERROR.2 :", err)
	return nil, err
}

// Create takes in user information and creates a new FIR user
func (us *UserServiceObj) Create(user *UserRequest) error {
	req, err := us.client.NewRequest("POST", usersPath, user)
	if err != nil {
		return err
	}

	resp, err := us.client.Do(req)

	if resp.StatusCode != 200 || err != nil {
		log.Println("Status Code:", resp.StatusCode)
		return err
	}

	return nil
}
