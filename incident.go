package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ListIncidents current FIR incidents
func ListIncidents(client *Client) (map[string]interface{}, error) {
	path := "/incidents"

	req, err := client.NewRequest("GET", path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := client.Do(req)

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

	fmt.Println("[ ERROR ] ", err)
	return nil, err

}
