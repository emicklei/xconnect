package xconnect

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestSpec(t *testing.T) {
	d, err := ioutil.ReadFile("spec-xconnect.yaml")
	if err != nil {
		t.Fatal(err)
	}
	type SpecHolder struct {
		X Config `yaml:"xconnect"`
	}
	var s Document
	if err := yaml.Unmarshal(d, &s); err != nil {
		t.Fatal(err)
	}
	x := s.Configuration
	if got, want := len(x.Listen), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(x.Connect), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// assert it all
	idc := x.Connect["<id>"]
	if got, want := idc.Scheme, "grpc"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := *idc.TLS, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := idc.Disabled, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := idc.Host, "there.com"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

}

func TestDumpSpec(t *testing.T) {
	x := Config{}
	x.Listen = map[string]ListenEntry{
		"api": ListenEntry{Scheme: "grpc"},
	}
	gcp := new(GCPEntry)
	gcp.Pubsub = new(GCPPubSubEntry)
	gcp.Pubsub.Topic = "topic"
	x.Connect = map[string]ConnectEntry{
		"db":  ConnectEntry{Scheme: "jdbc"},
		"sms": ConnectEntry{GCP: gcp},
	}
	data, err := yaml.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
