package firGo

// ArtifactService is an interface for interfacing with the Account
// endpoints of the DigitalOcean API
// See: https://developers.digitalocean.com/documentation/v2/#account
type ArtifactService interface {
	Get() (*Artifact, *Response, error)
}

// ArtifactServiceOp handles communication with the Account related methods of
// the DigitalOcean API.
type ArtifactServiceOp struct {
	client *Client
}

var _ ArtifactService = &ArtifactServiceOp{}

// Artifact represents a FIR Artifact
type Artifact struct {
	DropletLimit    int    `json:"droplet_limit,omitempty"`
	FloatingIPLimit int    `json:"floating_ip_limit,omitempty"`
	Email           string `json:"email,omitempty"`
	UUID            string `json:"uuid,omitempty"`
	EmailVerified   bool   `json:"email_verified,omitempty"`
	Status          string `json:"status,omitempty"`
	StatusMessage   string `json:"status_message,omitempty"`
}

type artifactRoot struct {
	Artifact *Artifact `json:"artifact"`
}

// Get FIR artifacts info
func (s *ArtifactServiceOp) Get() (*Artifact, *Response, error) {
	path := "/artifacts"

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(artifactRoot)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Artifact, resp, err
}
