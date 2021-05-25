package main

import (
	"os"
	"log"
	"encoding/json"
	
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph/parser"
)

func main() {
    configFile, err := os.Open("cmd/load_pipeline/pipeline.json")
    if err != nil {
        log.Println("opening config file", err.Error())
    }

	graphData := graph.GraphData{}
    jsonParser := json.NewDecoder(configFile)
    if err = jsonParser.Decode(&graphData); err != nil {
        log.Println("parsing config file", err.Error())
    }

	graph, err := parser.BuildGraph(graphData)
	if err != nil {
		log.Println("Can't build pipeline graph", err.Error())
	}

    log.Println(graph.Entrypoint.Data())

    return
}