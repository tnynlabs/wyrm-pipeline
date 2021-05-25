package nodes

import (
	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

type conditionNode struct {
	graph *g.Graph
	data *g.NodeData
}

// TODO: Node attributes validation

func BuildConditionNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	node := conditionNode{
		graph: graph,
		data: data,
	}

	return &node, nil
}

func (n *conditionNode) Run(input string) error {
	return nil
}

func (n *conditionNode) Data() g.NodeData {
	return *n.data
}

func (n *conditionNode) Type() g.NodeType {
	return g.ConditionNode
}
