package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type IncidentInterface interface {
	List() (map[string]interface{}, error)
	Create(map[string]interface{}) (Incident, error)
}

type Incident struct {
	Id 							int `json:"id",omitempty`
	Detection 			int `json:"detection",omitempty`
	Actor 					int `json:"actor",omitempty`
	Plan 						int `json:"plan",omitempty`
	FileSet 				[]string `json:"file_set",omitempty`
	Date 						string `json:"date",omitempty`
	IsStarred 			bool `json:"is_starred",omitempty`
	Subject 				string `json:"subject",omitempty`
	Description 		string `json:"description",omitempty`
	Severity 				int `json:"severity",omitempty`
	IsIncident		 	bool `json:"is_incident",omitempty`
	IsMajor	 				bool `json:"is_major",omitempty`
	Status 					string `json:"status",omitempty`
	Confidentiality int `json:"confidentiality",omitempty`
	Category 				int `json:"category",omitempty`
	OpenedBy 				int `json:"opened_by",omitempty`
	BizLines 				[]string `json:"concerned_business_lines",omitempty`
}

type IncidentServiceObj struct {
	client *Client
}

var _ IncidentInterface = &IncidentServiceObj{}
const incidentsPath = "/incidents"

func (a Incident) String() string {
	return Stringify(a)
}

// ListIncidents current FIR incidents
func (is *IncidentServiceObj) List() (map[string]interface{}, error) {
	req, err := is.client.NewRequest("GET", incidentsPath, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp, err := is.client.Do(req)

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

	fmt.Println("ERROR.2 ", err)
	return nil, err
}

func (is *IncidentServiceObj) Create(object map[string]interface{}) (Incident, error) {
	req, err := is.client.NewRequest("POST", incidentsPath, object)
	if err != nil {
		fmt.Println(err)
		return Incident{}, err
	}

	resp, err := is.client.Do(req)
	if err != nil {
		fmt.Println("ERROR.1: ", err)
	}
	fmt.Println(resp)
	return Incident{}, err
}
