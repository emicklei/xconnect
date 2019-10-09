package xconnect

import "log"

// GCPEntry is a Google Cloud Platform Service entry in connect.
type GCPEntry struct {
	Pubsub      *GCPPubSubEntry      `yaml:"pubsub,omitempty" json:"pubsub,omitempty"`
	MemoryStore *GCPMemoryStoreEntry `yaml:"memorystore,omitempty" json:"memorystore,omitempty"`
	DataStore   *GCPDataStoreEntry   `yaml:"datastore,omitempty" json:"datastore,omitempty"`
}

// GCPPubSubEntry is a Google Cloud Platform PubSub entry in gcp.
type GCPPubSubEntry struct {
	Subscription string `yaml:"subscription,omitempty" json:"subscription,omitempty"`
	Topic        string `yaml:"topic,omitempty" json:"topic,omitempty"`
}

// GCPMemoryStoreEntry is a Google Cloud Platform MemoryStore entry in gcp.
type GCPMemoryStoreEntry struct {
	InstanceID string `yaml:"instance-id" json:"instance-id"`
}

// GCPDataStoreEntry is a Google Cloud Platform Datastore entry in gcp.
type GCPDataStoreEntry struct {
	Namespace string `yaml:"namespace" json:"namespace"`
}

// PubSub finds the connect entry in the configuration. Panic with log if it does not.
func (c Config) PubSub(id string) GCPPubSubEntry {
	ce, ok := c.Connect[id]
	if !ok {
		log.Fatalf("no such connect entry [%s]", id)
	}
	gcp := ce.GCP
	if gcp == nil {
		log.Fatalf("no such connect gcp entry [%s]", id)
	}
	e := gcp.Pubsub
	if e == nil {
		log.Fatalf("no such connect gcp.pubsub entry [%s]", id)
	}
	return *e
}

// MemoryStore finds the connect entry in the configuration. Panic with log if it does not.
func (c Config) MemoryStore(id string) GCPMemoryStoreEntry {
	ce, ok := c.Connect[id]
	if !ok {
		log.Fatalf("no such connect entry [%s]", id)
	}
	gcp := ce.GCP
	if gcp == nil {
		log.Fatalf("no such connect gcp entry [%s]", id)
	}
	e := gcp.MemoryStore
	if e == nil {
		log.Fatalf("no such connect gcp.memorystore entry [%s]", id)
	}
	return *e
}

// DataStore finds the connect entry in the configuration. Panic with log if it does not.
func (c Config) DataStore(id string) GCPDataStoreEntry {
	ce, ok := c.Connect[id]
	if !ok {
		log.Fatalf("no such connect entry [%s]", id)
	}
	gcp := ce.GCP
	if gcp == nil {
		log.Fatalf("no such connect gcp entry [%s]", id)
	}
	e := gcp.DataStore
	if e == nil {
		log.Fatalf("no such connect gcp.datastore entry [%s]", id)
	}
	return *e
}
