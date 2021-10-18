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
	x := doc.XConnect
	if got, want := x.Connect["accounts"].FindString("gcp.datastore/kind"), "Account"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
