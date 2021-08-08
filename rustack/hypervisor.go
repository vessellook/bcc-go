package rustack

import "fmt"

type Hypervisor struct {
	manager *Manager
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

func (p *Project) GetAvailableHypervisors() (hypervisors []*Hypervisor, err error) {
	type tempType struct {
		Client struct {
			AllowedHypervisors []*Hypervisor `json:"allowed_hypervisors"`
		} `json:"client"`
	}

	var target tempType

	path := fmt.Sprintf("v1/project/%s", p.ID)
	err = p.manager.Get(path, Defaults(), &target)
	hypervisors = target.Client.AllowedHypervisors

	for i := range hypervisors {
		hypervisors[i].manager = p.manager
	}
	return
}
