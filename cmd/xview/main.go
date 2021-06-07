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

// xview | dot -Tpng  > graph.png && open graph.png

func main() {
	cfgs := []xconnect.Config{}
	for _, each := range collectYAMLnames() {
		d, err := loadDocument(each)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println("loaded", d.Config.Meta.Name)
		cfgs = append(cfgs, d.Config)
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

func addToGraph(cfg xconnect.Config, g *dot.Graph) {
	s := g.Subgraph(cfg.Meta.Name, dot.ClusterOption{})
	for k, v := range cfg.Listen {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		n := s.Node(id).Label(k)
		networkIDtoNode[v.NetworkID()] = n
	}
	for k := range cfg.Connect {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		// https://graphviz.org/doc/info/shapes.html#polygon
		s.Node(id).Label(k).Attr("shape", "box")
	}

}

func connectInGraph(cfg xconnect.Config, g *dot.Graph) {
	s, _ := g.FindSubgraph(cfg.Meta.Name)
	for k, v := range cfg.Connect {
		id := fmt.Sprintf("%s/%s", cfg.Meta.Name, k)
		from := s.Node(id)
		to, ok := networkIDtoNode[v.NetworkID()]
		if ok {
			from.Edge(to)
		} else {
			//fmt.Println("not found:", v.NetworkID())
		}
	}
}
