package firGo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const artifactsPath = "/artifacts"

// ArtifactInterface holds methods for the Artifact object
type ArtifactInterface interface {
	List() ([]Artifact, error)
	Create(*ArtifactRequest) error
}

// Artifact is an object model for FIR artifacts
type Artifact struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ArtifactRequest type is Artifact without an ID
type ArtifactRequest struct {
	Type  string `json:"type",omitempty`
	Value string `json:"value",omitempty`
}

// ArtifactResponse holds a response from FIR
type ArtifactResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []Artifact
}

// ArtifactServiceObj registers Artifact as a service available to the client
type ArtifactServiceObj struct {
	client *Client
}

var _ ArtifactInterface = &ArtifactServiceObj{}

// List lists FIR artifacts
func (as *ArtifactServiceObj) List() ([]Artifact, error) {
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

		var dat ArtifactResponse
		if err := json.Unmarshal(bodyBytes, &dat); err != nil {
			panic(err)
		}

		return dat.Results, nil
	}

	fmt.Println("ERROR.2 :", err)
	return nil, err
}

// Create adds a new artifact to FIR
func (as *ArtifactServiceObj) Create(a *ArtifactRequest) error {
	req, err := as.client.NewRequest("POST", artifactsPath, a)
	if err != nil {
		return err
	}

	resp, err := as.client.Do(req)

	if resp.StatusCode != 200 || err != nil {
		log.Println("Status Code:", resp.StatusCode)
		return err
	}
	return nil
}
