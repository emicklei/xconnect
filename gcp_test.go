package xconnect

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConnectToGCP(t *testing.T) {
	cfg := `
xconnect:	
  connect:
    accounts:
      gcp.datastore:
        kind: Account
`
	var doc Document
	if err := yaml.Unmarshal([]byte(cfg), &doc); err != nil {
		t.Fatal(err)
	}
	if got, want := doc.FindString("xconnect/connect/accounts/gcp.datastore/kind"), "Account"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
