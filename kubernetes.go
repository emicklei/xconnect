package xconnect

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

/**
apiVersion: v1
data:
	application.yml: |
		# comment
    xconnect:
		...
kind: ConfigMap
metadata:
  creationTimestamp: null
  name: envy
  namespace: envy
**/

// K8SConfiguration represents a Kubernetes configuration.
type K8SConfiguration struct {
	APIVersion string                 `yaml:"apiVersion"`
	Data       map[string]interface{} `yaml:"data"`
	Kind       string                 `yaml:"kind" `
	Metadata   struct {
		Name              string    `yaml:"name" `
		Namespace         string    `yaml:"namespace"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
	} `yaml:"metadata"`
}

// ExtractConfig expects a "xconnect" key in the data map and parses that part into a xconnect.Config.
func (k K8SConfiguration) ExtractConfig() (x XConnect, err error) {
	appYaml, ok := k.Data["application.yml"]
	if !ok {
		return x, errors.New("missing key: [application.yml]")
	}
	xconnectString, ok := appYaml.(string)
	if !ok {
		return x, fmt.Errorf("value of [application yml] is not a string] but [%T]", appYaml)
	}
	xconnectMap := map[string]interface{}{}
	if err = yaml.Unmarshal([]byte(xconnectString), &xconnectMap); err != nil {
		return x, err
	}
	xconnect, ok := xconnectMap["xconnect"]
	if !ok {
		return x, errors.New("missing key: [xconnect]")
	}
	// encode and decode again
	encoded, _ := yaml.Marshal(xconnect)
	err = yaml.Unmarshal(encoded, &x)
	return
}
