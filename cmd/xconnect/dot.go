package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/dot"
	"github.com/emicklei/xconnect"
	"gopkg.in/yaml.v2"
)

// read all xconnect config files
// start a webservice to display dot graphs in PNG

var master = dot.NewGraph(dot.Directed)
var networkIDtoNode = map[string]dot.Node{}

// xconnect -dot | dot -Tpng  > graph.png && open graph.png

func makeGraph() {
	cfgs := []xconnect.XConnect{}
	for _, each := range collectYAMLnames() {
		d, err := loadDocument(each)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println("loaded", d.Config.Meta.Name)
		cfgs = append(cfgs, d.XConnect)
	}
	for _, each := range cfgs {
		addToGraph(each, master)
	}
	for _, each := range cfgs {
		connectInGraph(each, master)
	}
	fmt.Println(master.String())
}

func collectYAMLnames() (files []string) {
	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yaml") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

func loadDocument(name string) (xconnect.Document, error) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return xconnect.Document{}, err
	}
	var d xconnect.Document
	if err = yaml.Unmarshal(content, &d); err != nil {
		return xconnect.Document{}, err
	}
	return d, nil
}

func addToGraph(cfg xconnect.XConnect, g *dot.Graph) {
	s := g.Subgraph(cfg.Meta.Name, dot.ClusterOption{})
	s.Attr("style", "rounded")
	s.Attr("bgcolor", "#F5FDF2")
	if bg, ok := cfg.Meta.ExtraFields["ui-bgcolor"]; ok {
		s.Attr("bgcolor", bg)
	}
	for k, v := range cfg.Listen {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		n := s.Node(id).Label(k)
		n.Attr("fillcolor", "#FFFFFF").Attr("style", "filled")
		if bg, ok := v.ExtraFields["ui-fillcolor"]; ok {
			n.Attr("fillcolor", bg).Attr("style", "filled")
		}
		networkIDtoNode[v.NetworkID()] = n
	}
	for k := range cfg.Connect {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		// https://graphviz.org/doc/info/shapes.html#polygon
		s.Node(id).Label(k).Attr("shape", "plaintext")
	}

}

func connectInGraph(cfg xconnect.XConnect, g *dot.Graph) {
	s, _ := g.FindSubgraph(cfg.Meta.Name)
	for k, v := range cfg.Connect {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		from := s.Node(id)
		to, ok := networkIDtoNode[v.NetworkID()]
		if !ok {
			// if kind is set then create a node to represent the other end
			if v.Kind != "" {
				id := v.ResourceID()
				to = g.Node(id)
				// remember
				networkIDtoNode[id] = to
			} else {
				fmt.Fprintf(os.Stderr, "[xconnect] no listen entry found: %s\n", v.NetworkID())
				continue
			}
		}
		from.Edge(to).Attr("arrowtail", "dot").Attr("dir", "both")
	}
}
