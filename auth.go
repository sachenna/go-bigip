package bigip

const (
	uriAuth        = "auth"
	uriTMPartition = "tmPartition"
)

// TMPartitions contains a list of all partitions on the BIG-IP system.
type TMPartitions struct {
	TMPartitions []*TMPartition `json:"items"`
}

type TMPartition struct {
	Name               string `json:"name,omitempty"`
	Kind               string `json:"kind,omitempty"`
	DefaultRouteDomain int    `json:"defaultRouteDomain,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
}

// TMPartitions returns a list of partitions.
func (b *BigIP) TMPartitions() (*TMPartitions, error) {
	var pList TMPartitions
	if err, _ := b.getForEntity(&pList, uriAuth, uriTMPartition); err != nil {
		return nil, err
	}
	return &pList, nil
}

// CreateTMPartitin creates a partition.
func (b *BigIP) CreateTMPartition(name string, routeDomain int) error {
	config := &TMPartition{
		Name:               name,
		DefaultRouteDomain: routeDomain,
	}

	return b.post(config, uriAuth, uriTMPartition)
}

// DeleteTMPartition removes a partition.
func (b *BigIP) DeleteTMPartition(name string) error {
	return b.delete(uriAuth, uriTMPartition, name)
}
