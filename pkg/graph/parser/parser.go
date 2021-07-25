package parser

import (
	"errors"
	"fmt"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"github.com/tnynlabs/wyrm-pipeline/pkg/graph/nodes"
)

func BuildGraph(data g.GraphData) (*g.Graph, error) {
	graph := g.Graph{
		Nodes: make(map[string]g.Node),
	}

	for _, nodeData := range data.Nodes {
		if _, exists := graph.Nodes[nodeData.Id]; exists {
			return nil, fmt.Errorf("duplicate node id '%s'", nodeData.Id)
		}

		node, err := BuildNode(&graph, nodeData)
		if err != nil {
			err = fmt.Errorf("node %s: %w", nodeData.Id, err)
			return nil, err
		}
		graph.Nodes[nodeData.Id] = node
	}

	entrypoint, ok := graph.Nodes[data.Entrypoint]
	if !ok {
		return nil, fmt.Errorf(
			"entrypoint node id '%s' not found", data.Entrypoint)
	}
	graph.Entrypoint = entrypoint

	// Todo: Is DAG graph validation
	// Todo: All children exist validation

	return &graph, nil
}

func BuildNode(graph *g.Graph, data g.NodeData) (g.Node, error) {
	switch data.Type {
	case g.ConditionNode:
		return nodes.BuildConditionNode(graph, &data)
	case g.DeviceInvokeNode:
		return nodes.BuildDeviceInvokeNode(graph, &data)
	case g.WebhookTriggerNode:
		return nodes.BuildWebhookTriggerNode(graph, &data)
	case g.FloatCastNode:
		return nodes.BuildFloatCastNode(graph, &data)
	case g.EmailNode:
		return nodes.BuildEmailNode(graph, &data)
	default:
		return nil, errors.New("invalid node type")
	}
}
