package parser

import (
	"fmt"
	"errors"
	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph/nodes"
)

func BuildGraph(data g.GraphData) (*g.Graph, error) {
	graph := g.Graph{
		Nodes: make(map[string]g.Node),
	}

	for _, nodeData := range data.Nodes {
		node, err := BuildNode(&graph, nodeData)
		if err != nil {
			return nil, err
		}
		if _, exists := graph.Nodes[nodeData.Id]; exists {
			return nil, fmt.Errorf("Duplicate node id '%s'", nodeData.Id)
		}
		graph.Nodes[nodeData.Id] = node
	}

	entrypoint, ok := graph.Nodes[data.Entrypoint]
	if !ok {
		return nil, fmt.Errorf(
			"Entrypoint node id '%s' not found", data.Entrypoint)
	}
	graph.Entrypoint = entrypoint

	// Todo: Is DAG graph validation
	// Todo: All children exist validation

	return &graph, nil
}

func BuildNode(graph *g.Graph, data g.NodeData) (g.Node, error) {
	switch data.Type {
	case g.DeviceResponseNode:
		return nodes.BuildDeviceResponseNode(graph, &data)
	case g.ConditionNode:
		return nodes.BuildConditionNode(graph, &data)
	case g.DeviceInvokeNode:
		return nodes.BuildDeviceInvokeNode(graph, &data)
	default:
		return nil, errors.New("Invalid node type")
	}
}
