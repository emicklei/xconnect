package xconnect

import "fmt"

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

func (e ListenEntry) find(keys []string) (interface{}, bool) {
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
	default:
		return findInMap(keys, e.ExtraFields)
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
