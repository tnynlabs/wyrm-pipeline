package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph/parser"
)

func main() {
	configFile, err := os.Open("cmd/load_pipeline/pipeline.json")
	if err != nil {
		log.Println("Opening config file", err.Error())
	}

	graphData := graph.GraphData{}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&graphData); err != nil {
		log.Println("Parsing config file", err.Error())
	}

	graph, err := parser.BuildGraph(graphData)
	if err != nil {
		log.Printf("Can't build pipeline graph (%v)\n", err.Error())
		return
	}

	log.Println(graph.Entrypoint.Data())
	err = graph.Entrypoint.Run(map[string]interface{}{
		"payload": map[string]interface{}{
			"name": "zozo",
			"age":  20,
		},
	})

	if err != nil {
		log.Println(err)
	}
}
