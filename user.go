package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// UsersInterface is an interface for all user struct methods
type UsersInterface interface {
	List() ([]User, error)
	Create(*User) error
}

// User type
type User struct {
	ID       int    `json:"id"`
	Groups   []int  `json:"groups"`
	Email    string `json:"email"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

// UserResponse holds a response from FIR
type UserResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []User
}

// UserServiceObj is a struct that allows client to
// receive new methods
type UserServiceObj struct {
	client *Client
}

const usersPath = "/users"

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
func (us *UserServiceObj) Create(user *User) error {
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
