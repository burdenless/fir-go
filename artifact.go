package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ListArtifacts lists FIR artifacts, returns a map
func ListArtifacts(client *Client) (map[string]interface{}, error) {
	path := "/artifacts"

	req, err := client.NewRequest("GET", path, nil)
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

	fmt.Println("ERROR.2 :", err)
	return nil, err
}
