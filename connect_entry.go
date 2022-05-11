package xconnect

import "fmt"

// ConnectEntry is a list element in the xconnect.connect config.
type ConnectEntry struct {
	Protocol string `yaml:"protocol,omitempty" json:"scheme,omitempty"`
	Secure   *bool  `yaml:"secure,omitempty" json:"secure,omitempty"`
	Host     string `yaml:"host,omitempty" json:"host,omitempty"`
	Port     *int   `yaml:"port,omitempty" json:"port,omitempty"`
	URL      string `yaml:"url,omitempty" json:"url,omitempty"`
	Disabled bool   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
	Kind     string `yaml:"kind,omitempty" json:"kind,omitempty"`
	// Resource is to identify the virtual listen part
	Resource    string                 `yaml:"resource,omitempty" json:"resource,omitempty"`
	ExtraFields map[string]interface{} `yaml:"-,inline"`
}

type ConnectionEnd interface {
	NetworkID() string
}

// ResourceID returns NetworkID() or KIND:RESOURCE
func (e ConnectEntry) ResourceID() string {
	if id := e.NetworkID(); id != "" {
		return id
	}
	return fmt.Sprintf("%s:%s", e.Kind, e.Resource)
}

// NetworkID returns URL or HOST:PORT
func (e ConnectEntry) NetworkID() string {
	// URL overrides Host+Port
	if len(e.URL) != 0 {
		return e.URL
	}
	if len(e.Host) != 0 {
		p := 0
		if e.Port != nil {
			p = *e.Port
		}
		return fmt.Sprintf("%s:%d", e.Host, p)
	}
	// url empty, host empty, we dont know
	return ""
}

func (e ConnectEntry) find(keys []string) (interface{}, bool) {
	if len(keys) == 0 {
		return nil, false
	}
	switch keys[0] {
	case "protocol":
		return e.Protocol, true
	case "secure":
		if e.Secure != nil {
			return *e.Secure, true
		}
		return nil, false
	case "host":
		return e.Host, true
	case "port":
		if e.Port != nil {
			return *e.Port, true
		}
		return nil, false
	case "url":
		return e.URL, true
	case "disabled":
		return e.Disabled, true
	case "kind":
		return e.Kind, true
	case "resource":
		return e.Resource, true
	default:
		return findInMap(keys, e.ExtraFields)
	}
}
