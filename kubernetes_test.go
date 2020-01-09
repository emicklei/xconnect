package xconnect

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestKubernetesSome(t *testing.T) {
	d, err := ioutil.ReadFile("kubernetes_configmap-application.properties.yml")
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
	if got, want := len(x.Connect["variant-publish"].ExtraFields), 1; got != want {
		t.Fatalf("got [%d] extra fields want [%d]", got, want)
	}
	if got, want := x.Connect["variant-publish"].FindString("gcp.pubsub/topic"), "VariantToAssortment_Push_v1-topic"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
	if got, want := len(x.Connect["variant-pull"].ExtraFields), 1; got != want {
		t.Fatalf("got [%d] extra fields want [%d]", got, want)
	}
	if got, want := x.Connect["variant-pull"].FindString("gcp.pubsub/subscription"), "Variant_v1-subscription"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
