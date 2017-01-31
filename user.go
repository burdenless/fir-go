package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UserInterface interface {
	List() (map[string]interface{}, error)
}

type User struct {
	Id 					int `json:"id"`
	Groups 			[]int `json:"groups"`
	Email 			string `json:"email"`
	Username 		string `json:"username"`
	URL 				string `json:"url"`
}

type UserServiceObj struct {
	client *Client
}

const usersPath = "/users"

// ListUsers current FIR incidents
func (us *UserServiceObj) List() (map[string]interface{}, error) {
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

		var dat map[string]interface{}
    if err := json.Unmarshal(bodyBytes, &dat); err != nil {
        panic(err)
    }

		return dat, nil
	}

	fmt.Println("ERROR.2 :", err)
	return nil, err
}

func (us *UserServiceObj) Create(object map[string]interface{}) (User, error) {
	_, err := us.client.NewRequest("POST", usersPath, object)

	return User{}, err
}
