package nodes

import (
	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

type deviceResponseNode struct {
	graph *g.Graph
	data *g.NodeData
}

func BuildDeviceResponseNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	node := deviceResponseNode{
		graph: graph,
		data: data,
	}

	return &node, nil
}

func (n *deviceResponseNode) Run(input string) error {
	return nil
}

func (n *deviceResponseNode) Data() g.NodeData {
	return *n.data
}

func (n *deviceResponseNode) Type() g.NodeType {
	return g.DeviceResponseNode
}
