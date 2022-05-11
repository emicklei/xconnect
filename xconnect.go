package xconnect

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const extraPathSeparator = "/"

// XConnect represents the xconnect data section of a YAML document.
// See spec-xconnect.yaml.
type XConnect struct {
	Meta        MetaProperties          `yaml:"meta" json:"meta"`
	Listen      map[string]ListenEntry  `yaml:"listen" json:"listen"`
	Connect     map[string]ConnectEntry `yaml:"connect" json:"connect"`
	ExtraFields map[string]interface{}  `yaml:"-,inline"`
}

func (x XConnect) find(keys []string) (interface{}, bool) {
	if len(keys) == 0 {
		return nil, false
	}
	switch keys[0] {
	case "meta":
		return x.Meta.find(keys[1:])
	case "listen":
		subkeys := keys[1:]
		if len(subkeys) == 0 {
			return nil, false
		}
		for k, each := range x.Listen {
			if subkeys[0] == k {
				if v, ok := each.find(subkeys[1:]); ok {
					return v, ok
				}
			}
		}
		return nil, false
	case "connect":
		subkeys := keys[1:]
		if len(subkeys) == 0 {
			return nil, false
		}
		for k, each := range x.Connect {
			if subkeys[0] == k {
				if v, ok := each.find(subkeys[1:]); ok {
					return v, ok
				}
			}
		}
		return nil, false
	default:
		return findInMap(keys, x.ExtraFields)
	}
}

// MetaProperties represents the meta element in the xconnect data section.
type MetaProperties struct {
	Name    string `yaml:"name,omitempty" json:"name,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
	// Operational expenditure, or owner
	Opex        string                 `yaml:"opex,omitempty" json:"opex,omitempty"`
	Labels      []string               `yaml:"tags,omitempty" json:"tags,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
	Kind        string                 `yaml:"kind,omitempty" json:"kind,omitempty"`
}

func (m MetaProperties) find(keys []string) (interface{}, bool) {
	if len(keys) == 0 {
		return nil, false
	}
	switch keys[0] {
	case "name":
		return m.Name, true
	case "version":
		return m.Version, true
	case "opex":
		return m.Opex, true
	case "kind":
		return m.Kind, true
	default:
		return findInMap(keys, m.ExtraFields)
	}
}

// Document is the root YAML element
type Document struct {
	XConnect    XConnect               `yaml:"xconnect"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

// MustString same as FindString but panics if not found. E.g xconnect/connect/db/url .
func (d Document) MustString(path string) string {
	if v, err := d.FindString(path); err != nil {
		panic(err)
	} else {
		return v
	}
}

// FindString returns a string for a given slash path, e.g xconnect/connect/db/url .
func (d Document) FindString(path string) (string, error) {
	return FindString(d, path)
}

// FindString returns a string for a given slash path, e.g xconnect/connect/db/url .
func FindString(f finder, path string) (string, error) {
	keys := strings.Split(path, extraPathSeparator)
	v, ok := f.find(keys)
	if !ok {
		return "", fmt.Errorf("unable to find string at [%s]", path)
	}
	if s, ok := v.(string); !ok {
		return "", fmt.Errorf("warn: xconnect, value is not a string, but a %T for path %s\n", v, path)
	} else {
		return s, nil
	}
}

// MustBool same as FindBool but panics if not found. E.g xconnect/listen/api/secure
func (d Document) MustBool(path string) bool {
	if v, err := d.FindBool(path); err != nil {
		panic(err)
	} else {
		return v
	}
}

// FindBool returns a bool for a given slash path.
func (d Document) FindBool(path string) (bool, error) {
	return FindBool(d, path)
}

func FindBool(f finder, path string) (bool, error) {
	keys := strings.Split(path, extraPathSeparator)
	v, ok := f.find(keys)
	if !ok {
		return false, fmt.Errorf("unable to find bool at [%s]", path)
	}
	if s, ok := v.(bool); !ok {
		return false, fmt.Errorf("warn: xconnect, value is not a bool, but a %T for path %s\n", v, path)
	} else {
		return s, nil
	}
}

// MustInt same as FindInt but panics if not found. E.g xconnect/listen/api/port
func (d Document) MustInt(path string) int {
	if v, err := d.FindInt(path); err != nil {
		panic(err)
	} else {
		return v
	}
}

// FindInt returns a integer for a given slash path, e.g xconnect/listen/api/port .
func (d Document) FindInt(path string) (int, error) {
	return FindInt(d, path)
}

// FindInt returns a integer for a given slash path.
func FindInt(f finder, path string) (int, error) {
	keys := strings.Split(path, extraPathSeparator)
	v, ok := f.find(keys)
	if !ok {
		return 0, fmt.Errorf("unable to find int at [%s]", path)
	}
	if s, ok := v.(int); !ok {
		return 0, fmt.Errorf("warn: xconnect, value is not a int, but a %T for path %s\n", v, path)
	} else {
		return s, nil
	}
}

func (d Document) find(keys []string) (interface{}, bool) {
	if len(keys) == 0 {
		return nil, false
	}
	switch keys[0] {
	case "xconnect":
		return d.XConnect.find(keys[1:])
	default:
		return findInMap(keys, d.ExtraFields)
	}
}

// GetConfig will first check the environment value at {envKey} to find the source of the confguration.
// If the environment value is not available (empty) then try reading the filename to get the configuration.
func GetConfig(envKey string, filename string) (Document, error) {
	content := os.Getenv(envKey)
	if len(content) == 0 {
		return LoadConfig(filename)
	}
	var doc Document
	err := yaml.Unmarshal([]byte(content), &doc)
	if err != nil {
		return Document{}, fmt.Errorf("unable to unmarshal YAML:%v", err)
	}
	return doc, nil
}

// LoadConfig returns the document containing the xconnect section.
func LoadConfig(filename string) (Document, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return Document{}, fmt.Errorf("unable to read:%v", err)
	}
	var doc Document
	err = yaml.Unmarshal(content, &doc)
	if err != nil {
		return Document{}, fmt.Errorf("unable to unmarshal YAML:%v", err)
	}
	return doc, nil
}
