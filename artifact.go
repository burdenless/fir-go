package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ArtifactInterface interface {
	List() (map[string]interface{}, error)
}

type Artifact struct {
	Id 				int 		 `json:"id"`
	Type 			string 	 `json:"type"`
	Value 		string 	 `json:"value"`
	Artifacts []string `json:"artifacts"`
}

type ArtifactServiceObj struct {
	client *Client
}

var _ ArtifactInterface = &ArtifactServiceObj{}
const artifactsPath = "/artifacts"

func (a Artifact) String() string {
	return Stringify(a)
}

// ListArtifacts lists FIR artifacts, returns a map
func (as *ArtifactServiceObj) List() (map[string]interface{}, error) {
	req, err := as.client.NewRequest("GET", artifactsPath, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := as.client.Do(req)

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
