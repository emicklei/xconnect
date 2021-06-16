package xconnect

import (
	"fmt"
	"log"
	"strings"
)

// ListenEntry is a list element in the xconnect.accept config.
type ListenEntry struct {
	Protocol string `yaml:"protocol,omitempty" json:"scheme,omitempty"`
	Host     string `yaml:"host,omitempty" json:"host,omitempty"`
	Port     *int   `yaml:"port,omitempty" json:"port,omitempty"`
	// for database connection strings
	URL         string                 `yaml:"url,omitempty" json:"url,omitempty"`
	Secure      *bool                  `yaml:"secure,omitempty" json:"secure,omitempty"`
	Disabled    bool                   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

const extraPathSeparator = "/"

// FindString return a string for a given path (using slashes).
func (e ListenEntry) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	withFixed["protocol"] = e.Protocol
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

func (e ListenEntry) NetworkID() string {
	// URL overrides Host+Port
	if len(e.URL) != 0 {
		return e.URL
	}
	p := 0
	if e.Port != nil {
		p = *e.Port
	}
	return fmt.Sprintf("%s:%d", e.Host, p)
}

// ConnectEntry is a list element in the xconnect.connect config.
type ConnectEntry struct {
	Protocol    string                 `yaml:"protocol,omitempty" json:"scheme,omitempty"`
	Secure      *bool                  `yaml:"secure,omitempty" json:"secure,omitempty"`
	Host        string                 `yaml:"host,omitempty" json:"host,omitempty"`
	Port        *int                   `yaml:"port,omitempty" json:"port,omitempty"`
	URL         string                 `yaml:"url,omitempty" json:"url,omitempty"`
	Disabled    bool                   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	Kind        string                 `yaml:"kind,omitempty" json:"kind,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

type ConnectionEnd interface {
	NetworkID() string
}

// FindString return a string for a give dotted path.
func (e ConnectEntry) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(e.ExtraFields)
	withFixed["protocol"] = e.Protocol
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

func (e ConnectEntry) NetworkID() string {
	// URL overrides Host+Port
	if len(e.URL) != 0 {
		return e.URL
	}
	p := 0
	if e.Port != nil {
		p = *e.Port
	}
	return fmt.Sprintf("%s:%d", e.Host, p)
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
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
	// Operational expenditure, or owner
	Opex        string                 `yaml:"opex,omitempty" json:"opex,omitempty"`
	Labels      []string               `yaml:"tags,omitempty" json:"tags,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

// FindString return a string for a given slash path.
func (m MetaConfig) FindString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	withFixed := copy(m.ExtraFields)
	withFixed["name"] = m.Name
	withFixed["version"] = m.Version
	withFixed["opex"] = m.Opex
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
