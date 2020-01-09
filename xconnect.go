package xconnect

import (
	"log"
	"strings"
)

// ListenEntry is a list element in the xconnect.accept config.
type ListenEntry struct {
	Scheme      string                 `yaml:"scheme,omitempty" json:"scheme,omitempty"`
	Host        string                 `yaml:"host,omitempty" json:"host,omitempty"`
	Port        *int                   `yaml:"port,omitempty" json:"port,omitempty"`
	TLS         *bool                  `yaml:"tls,omitempty" json:"tls,omitempty"`
	Disabled    bool                   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

const extraPathSeparator = "/"

// FindString return a string for a given path (using slashes).
func (e ListenEntry) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	withFixed["scheme"] = e.Scheme
	withFixed["host"] = e.Host
	v := find(keys, withFixed)
	if s, ok := v.(string); !ok {
		log.Printf("warn: xconnect, value is not a string, but a %T for path %s\n", v, path)
		return ""
	} else {
		return s
	}
}

// FindInt returns an int for a given path (using slashes).
func (e ListenEntry) FindInt(path string) int {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	if e.Port != nil {
		withFixed["port"] = *e.Port
	}
	v := find(keys, withFixed)
	if i, ok := v.(int); !ok {
		log.Printf("warn: xconnect, value is not a int, but a %T for path %s\n", v, path)
		return 0
	} else {
		return i
	}
}

// ConnectEntry is a list element in the xconnect.connect config.
type ConnectEntry struct {
	Scheme      string                 `yaml:"scheme,omitempty" json:"scheme,omitempty"`
	TLS         *bool                  `yaml:"tls,omitempty" json:"tls,omitempty"`
	Host        string                 `yaml:"host,omitempty" json:"host,omitempty"`
	Port        *int                   `yaml:"port,omitempty" json:"port,omitempty"`
	URL         string                 `yaml:"url,omitempty" json:"url,omitempty"`
	Disabled    bool                   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

// FindString return a string for a give dotted path.
func (e ConnectEntry) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	withFixed["scheme"] = e.Scheme
	withFixed["host"] = e.Host
	withFixed["url"] = e.URL
	v := find(keys, withFixed)
	if s, ok := v.(string); !ok {
		log.Printf("warn: xconnect, value is not a string, but a %T for path %s\n", v, path)
		return ""
	} else {
		return s
	}
}

// FindInt returns an int for a given path (using slashes).
func (e ConnectEntry) FindInt(path string) int {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	if e.Port != nil {
		withFixed["port"] = *e.Port
	}
	v := find(keys, withFixed)
	if i, ok := v.(int); !ok {
		log.Printf("warn: xconnect, value is not a int, but a %T for path %s\n", v, path)
		return 0
	} else {
		return i
	}
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
	Name        string                 `yaml:"name,omitempty" json:"name,omitempty"`
	Version     string                 `yaml:"version,omitempty" json:"version,omitempty"`
	Owner       string                 `yaml:"owner,omitempty" json:"owner,omitempty"`
	Labels      []string               `yaml:"labels,omitempty" json:"labels,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

// FindString return a string for a given slash path.
func (m MetaConfig) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(m.ExtraFields)
	withFixed["name"] = m.Name
	withFixed["version"] = m.Version
	withFixed["owner"] = m.Owner
	v := find(keys, withFixed)
	if s, ok := v.(string); !ok {
		log.Printf("warn: xconnect, value is not a string, but a %T for path %s\n", v, path)
		return ""
	} else {
		return s
	}
}

// Document is the root YAML element
type Document struct {
	Config Config `yaml:"xconnect"`
}
