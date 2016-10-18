package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Incident represents a FIR Incident
type Incident struct {
	Detection   int      `json:"detection,omitempty"`
	Actor       int      `json:"actor,omitempty"`
	Plan        int      `json:"plan,omitempty"`
	FileSet     []string `json:"file_set,omitempty"`
	Date        string   `json:"date,omitempty"`
	Subject     string   `json:"subject,omitempty"`
	Description string   `json:"description,omitempty"`
	Severity    int      `json:"severity,omitempty"`
	IsIncident  bool     `json:"is_incident,omitempty"`
	IsStarred   bool     `json:"is_starred,omitempty"`
	IsMajor     bool     `json:"is_major,omitempty"`
}

// IncidentList is an array of Incidents
type IncidentList []Incident

// ListIncidents current FIR incidents
func ListIncidents(client *Client) (IncidentList, error) {
	path := "/incidents"

	req, err := client.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		bodyString := string(bodyBytes)
		res := []Incident{}
		json.Unmarshal([]byte(bodyString), &res)
		fmt.Println("[ Server Response ]", resp)

		return res, nil
	}

	fmt.Println("[ ERROR ] ", err)
	return nil, err

}
