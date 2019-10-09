package xconnect

// ListenEntry is a list element in the xconnect.accept config.
type ListenEntry struct {
	Scheme   string `yaml:"scheme,omitempty" json:"scheme,omitempty"`
	TLS      *bool  `yaml:"tls,omitempty" json:"tls,omitempty"`
	Port     *int   `yaml:"port,omitempty" json:"port,omitempty"`
	Disabled bool   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
}

// ConnectEntry is a list element in the xconnect.connect config.
type ConnectEntry struct {
	Scheme   string `yaml:"scheme,omitempty" json:"scheme,omitempty"`
	TLS      *bool  `yaml:"tls,omitempty" json:"tls,omitempty"`
	Host     string `yaml:"host,omitempty" json:"host,omitempty"`
	Port     *int   `yaml:"port,omitempty" json:"port,omitempty"`
	URL      string `yaml:"url,omitempty" json:"url,omitempty"`
	Disabled bool   `yaml:"disabled,omitempty" json:"disabled,omitempty"`

	// GoogleServices
	GCP *GCPEntry `yaml:"gcp,omitempty" json:"gcp,omitempty"`
}

// Config represents the xconnect data section of a YAML document.
// See spec-xconnect.yaml.
type Config struct {
	Meta    MetaConfig              `yaml:"meta" json:"meta"`
	Listen  map[string]ListenEntry  `yaml:"listen" json:"listen"`
	Connect map[string]ConnectEntry `yaml:"connect" json:"connect"`
}

// MetaConfig represents the meta element in the xconnect data section.
type MetaConfig struct {
	Name    string   `yaml:"name,omitempty" json:"name,omitempty"`
	Version string   `yaml:"version,omitempty" json:"version,omitempty"`
	Owner   string   `yaml:"owner,omitempty" json:"owner,omitempty"`
	Labels  []string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

// Document is the root YAML element
type Document struct {
	Configuration Config `yaml:"xconnect"`
}
