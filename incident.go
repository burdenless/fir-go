package firGo

import "fmt"

// IncidentService is an interface for interfacing with the Account
// endpoints of the DigitalOcean API
// See: https://developers.digitalocean.com/documentation/v2/#account
type IncidentService interface {
	Get() (*Incident, *Response, error)
}

// IncidentServiceOp handles communication with the Incident related methods of
// the Fast Incident Response API.
type IncidentServiceOp struct {
	client *Client
}

var _ IncidentService = &IncidentServiceOp{}

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

// IncidentRoot is the root of the Incident object
type IncidentRoot struct {
	Incident *Incident `json:"incident"`
}

// Get FIR artifacts info
func (s *IncidentServiceOp) Get() (*Incident, *Response, error) {
	path := "/incidents"
	fmt.Println("Calling incidents endpoint!")
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(IncidentRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Incident, resp, err
}
