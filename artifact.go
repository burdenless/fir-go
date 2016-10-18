package firGo

import "fmt"

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

// ArtifactList lists FIR artifacts info
func ArtifactList(client *Client) error {
	path := "/artifacts"

	req, err := client.NewRequest("GET", path)
	if err != nil {
		return err
	}

	fmt.Println(req)

	return err
}
