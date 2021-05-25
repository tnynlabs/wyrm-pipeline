package graph

import (
	"errors"
)

type Graph struct {
	Entrypoint Node
	Nodes map[string]Node
}

func (g *Graph) GetNode(id string) (Node, error) {
	node, ok := g.Nodes[id]
	if !ok {
		return nil, errors.New("Node not found")
	}
	return node, nil
}

type GraphData struct {
	Entrypoint string	`json:"entrypoint"`
	Nodes []NodeData	`json:"nodes"`
}
