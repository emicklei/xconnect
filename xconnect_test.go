package xconnect

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExtended(t *testing.T) {
	d, err := ioutil.ReadFile("extended-xconnect.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc Document
	if err := yaml.Unmarshal(d, &doc); err != nil {
		t.Fatal(err)
	}
	c := doc.Config
	if got, want := c.Meta.ExtraString("extra0"), "extra0"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c.Meta.ExtraString("nested0.sub0"), "sub0"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c.Listen["id1"].ExtraString("extra1"), "extra1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c.Listen["id1"].ExtraString("nested1.sub1"), "sub1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c.Connect["id2"].ExtraString("extra2"), "extra2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c.Connect["id2"].ExtraString("nested2.sub2"), "sub2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestSpec(t *testing.T) {
	d, err := ioutil.ReadFile("spec-xconnect.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc Document
	if err := yaml.Unmarshal(d, &doc); err != nil {
		t.Fatal(err)
	}
	x := doc.Config
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
