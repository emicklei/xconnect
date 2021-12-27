package xconnect

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExtended(t *testing.T) {
	d, err := ioutil.ReadFile("xconnect-extended.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc Document
	if err := yaml.Unmarshal(d, &doc); err != nil {
		t.Fatal(err)
	}
	if got, want := doc.ExtraFields["any"], "value"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	s, _ := doc.FindString("any")
	if got, want := s, "value"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	i, _ := doc.FindInt("xconnect/int")
	if got, want := i, 2; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	s, _ = doc.FindString("xconnect/any")
	if got, want := s, "other"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	s, _ = doc.FindString("xconnect/meta/extra0")
	if got, want := s, "extra0"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/meta/nested0/sub0")
	if got, want := s, "sub0"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/listen/id1/extra1")
	if got, want := s, "extra1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/listen/id1/nested1/sub1")
	if got, want := s, "sub1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/connect/id2/extra2")
	if got, want := s, "extra2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/connect/id2/nested2/sub2")
	if got, want := s, "sub2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	s, _ = doc.FindString("xconnect/connect/id2/host")
	if got, want := s, "notextra"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	i, _ = doc.FindInt("xconnect/connect/id2/port")
	if got, want := i, -1; got != want {
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
	x := doc.XConnect
	if got, want := len(x.Listen), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(x.Connect), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// assert it all
	idc := x.Connect["<id>"]
	if got, want := idc.Protocol, "grpc"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := *idc.Secure, true; got != want {
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
	x := XConnect{}
	x.Listen = map[string]ListenEntry{
		"api": {Protocol: "grpc"},
	}
	x.Connect = map[string]ConnectEntry{
		"db": {Protocol: "jdbc"},
	}
	data, err := yaml.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig("spec-xconnect.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
