package nodes

import (
	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

type deviceInvokeNode struct {
	graph *g.Graph
	data *g.NodeData
}

func BuildDeviceInvokeNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	node := deviceInvokeNode{
		graph: graph,
		data: data,
	}

	return &node, nil
}

func (n *deviceInvokeNode) Run(input string) error {
	return nil
}

func (n *deviceInvokeNode) Data() g.NodeData {
	return *n.data
}

func (n *deviceInvokeNode) Type() g.NodeType {
	return g.DeviceInvokeNode
}
