package xconnect

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestSome(t *testing.T) {
	d, err := ioutil.ReadFile("some-configmap-application.properties.yml")
	if err != nil {
		t.Fatal(err)
	}
	var k K8SConfiguration
	if err := yaml.Unmarshal(d, &k); err != nil {
		t.Fatal(err)
	}
	x, err := k.ExtractConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(x.Listen), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(x.Connect), 4; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	if x.Connect["variant-pull"].GCP == nil {
		t.Fatal("missing variant-pull gcp", x.Connect["variant-pull"])
	}
	if got, want := x.Connect["variant-pull"].GCP.Pubsub.Subscription, "Variant_v1-subscription"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
