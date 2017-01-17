package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Incident struct {
		Id 							int `json:"id"`
		Detection 			int `json:"detection"`
		Actor 					int `json:"actor"`
		Plan 						int `json:"plan"`
		FileSet 				[]string `json:"file_set"`
		Date 						string `json:"date"`
		IsStarred 			bool `json:"is_starred"`
		Subject 				string `json:"subject"`
		Description 		string `json:"description"`
		Severity 				int `json:"severity"`
		IsIncident		 	bool `json:"is_incident"`
		IsMajor	 				bool `json:"is_major"`
		Status 					string `json:"status"`
		Confidentiality int `json:"confidentiality"`
		Category 				int `json:"category"`
		OpenedBy 				int `json:"opened_by"`
		BizLines 				[]string `json:"concerned_business_lines"`
}

const path = "/incidents"

// ListIncidents current FIR incidents
func ListIncidents(client *Client) (map[string]interface{}, error) {
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

	fmt.Println("[ ERROR ] ", err)
	return nil, err
}

func AddIncident(client *Client, object map[string]interface{}) (Incident, error) {
	req, err := client.NewRequest("POST", path, object)
	if err != nil {
		fmt.Println(err)
		return Incident{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR.1: ", err)
	}
	fmt.Println(resp)
	return Incident{}, err
}
