package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/emicklei/xconnect"
	"gopkg.in/yaml.v2"
)

var oInput = flag.String("input", "xconnect.yaml", "name of the YAML configuration file that contains a xconnect section")
var oK8S = flag.Bool("k8s", false, "YAML is a Kubernetes configuration file with data:xconnect section")
var oTarget = flag.String("target", "http://localhost:8080", "destination for the YAML representation of the xconnect configuration, http or file scheme")

func main() {
	flag.Parse()

	log.Println("READ ", *oInput)
	content, err := ioutil.ReadFile(*oInput)
	if err != nil {
		log.Fatal(err)
	}
	var cfg xconnect.Config
	if *oK8S { // get xconnect section from k8s configuration
		extracted, err := readK8S(content)
		if err != nil {
			log.Fatal(err)
		}
		cfg = extracted
	} else {
		extracted, err := readXConnectDocument(content)
		if err != nil {
			log.Fatal(err)
		}
		cfg = extracted
	}
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	if err := enc.Encode(cfg); err != nil {
		log.Fatal("unable to marshal into YAML", err)
	}
	if strings.HasPrefix(*oTarget, "http") {
		log.Println("POST", *oTarget)
		resp, err := http.Post(*oTarget, "image/jpeg", &buf)
		if err != nil {
			log.Fatal("unable to POST configuration", err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatal("unable to POST configuration", err)
		}
		return
	}
	if strings.HasPrefix(*oTarget, "file") {
		withoutScheme := (*oTarget)[len("file://"):]
		log.Println("WRITE", withoutScheme)
		err := ioutil.WriteFile(withoutScheme, buf.Bytes(), os.ModePerm)
		if err != nil {
			log.Fatal("unable to write configuration", err)
		}
		return
	}
}

func readXConnectDocument(content []byte) (cfg xconnect.Config, err error) {
	log.Println("PARSE xconnect configuration ", *oInput)
	var d xconnect.Document
	if err = yaml.Unmarshal(content, &d); err != nil {
		return
	}
	return d.Configuration, nil
}

func readK8S(content []byte) (cfg xconnect.Config, err error) {
	log.Println("PARSE Kubernetes (k8s) Configuration ", *oInput)
	var k xconnect.K8SConfiguration
	if err = yaml.Unmarshal(content, &k); err != nil {
		return
	}
	x, err := k.ExtractConfig()
	if err != nil {
		return
	}
	return x, nil
}
