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

// ExtraString return a string for a given path (using slashes).
func (e ListenEntry) ExtraString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	return findString(keys, e.ExtraFields)
}

func findString(path []string, tree map[string]interface{}) string {
	if len(tree) == 0 {
		log.Println("warn: xconnect, empty extra fields")
		return ""
	}
	if len(path) == 0 {
		log.Println("warn: xconnect, empty key", path[0])
		return ""
	}
	if len(path) == 1 {
		f, ok := tree[path[0]]
		if !ok {
			log.Println("warn: xconnect, no such key", path[0])
			return ""
		}
		s, ok := f.(string)
		if !ok {
			log.Println("warn: xconnect, value not a string for key", path[0])
			return ""
		}
		return s
	}
	// > 1
	f, ok := tree[path[0]]
	if !ok {
		log.Println("warn: xconnect, no such key", path[0])
		return ""
	}
	mi, ok := f.(map[interface{}]interface{})
	if !ok {
		log.Printf("warn: xconnect, value is not a map, but a %T for key %s\n", f, path[0])
		return ""
	}
	// do the copy
	m := map[string]interface{}{}
	for k, v := range mi {
		sk, ok := k.(string)
		if !ok {
			log.Printf("warn: xconnect, key %s is not a string but %T\n", k, k)
		} else {
			m[sk] = v
		}
	}
	return findString(path[1:], m)
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

// ExtraString return a string for a give dotted path.
func (e ConnectEntry) ExtraString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	return findString(keys, e.ExtraFields)
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

// ExtraString return a string for a give dotted path.
func (m MetaConfig) ExtraString(path string) string {
	keys := strings.Split(path, extraPathSeparator)
	return findString(keys, m.ExtraFields)
}

// Document is the root YAML element
type Document struct {
	Config Config `yaml:"xconnect"`
}
