package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const incidentsPath = "/incidents"

// IncidentInterface holds methods for Incident objects
type IncidentInterface interface {
	List() ([]Incident, error)
	Create(*IncidentRequest) error
}

// Incident struct is the model for FIR incidents
type Incident struct {
	ID              int      `json:"id",omitempty`
	Detection       int      `json:"detection",omitempty`
	Actor           int      `json:"actor",omitempty`
	Plan            int      `json:"plan",omitempty`
	FileSet         []string `json:"file_set",omitempty`
	Date            string   `json:"date",omitempty`
	IsStarred       bool     `json:"is_starred",omitempty`
	Subject         string   `json:"subject",omitempty`
	Description     string   `json:"description",omitempty`
	Severity        int      `json:"severity",omitempty`
	IsIncident      bool     `json:"is_incident",omitempty`
	IsMajor         bool     `json:"is_major",omitempty`
	Status          string   `json:"status",omitempty`
	Confidentiality int      `json:"confidentiality",omitempty`
	Category        int      `json:"category",omitempty`
	OpenedBy        int      `json:"opened_by",omitempty`
	BizLines        []int    `json:"concerned_business_lines",omitempty`
}

// IncidentRequest is a model for creating/updating incidents
type IncidentRequest struct {
	Detection       int    `json:"detection",omitempty`
	Actor           int    `json:"actor",omitempty`
	Plan            int    `json:"plan",omitempty`
	Date            string `json:"date",omitempty`
	IsStarred       bool   `json:"is_starred",omitempty`
	Subject         string `json:"subject",omitempty`
	Description     string `json:"description",omitempty`
	Severity        int    `json:"severity",omitempty`
	IsIncident      bool   `json:"is_incident",omitempty`
	IsMajor         bool   `json:"is_major",omitempty`
	Confidentiality int    `json:"confidentiality",omitempty`
	Category        int    `json:"category",omitempty`
	OpenedBy        int    `json:"opened_by",omitempty`
	//BizLines        []string `json:"concerned_business_lines",omitempty`
}

// IncidentServiceObj registers Incidents as a service to the client
type IncidentServiceObj struct {
	client *Client
}

// IncidentResponse holds the response from FIR with Incidents
type IncidentResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []Incident
}

var _ IncidentInterface = &IncidentServiceObj{}

// List current FIR incidents
func (is *IncidentServiceObj) List() ([]Incident, error) {
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

		var dat IncidentResponse
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}

		return dat.Results, nil
	}

	fmt.Println("ERROR.2 ", err)
	return nil, err
}

// Create will add a new incident object
func (is *IncidentServiceObj) Create(i *IncidentRequest) error {
	// Set some incident defaults if they are not present
	i.Date = time.Now().UTC().Format(time.RFC3339Nano)
	if i.Severity == 0 {
		i.Severity = 1
	}
	if i.Detection == 0 {
		i.Detection = 1
	}
	if i.Actor == 0 {
		i.Actor = 3
	}
	if i.Plan == 0 {
		i.Plan = 5
	}
	if i.Category == 0 {
		i.Category = 1
	}

	req, err := is.client.NewRequest("POST", incidentsPath, i)
	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, err := is.client.Do(req)
	if err != nil {
		fmt.Println("ERROR.1: ", err)
	}

	if resp.StatusCode != 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("Status code: %d, Error: %s", resp.StatusCode, bodyString)
	}
	return nil
}
