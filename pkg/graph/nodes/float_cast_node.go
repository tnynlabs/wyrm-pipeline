package nodes

import (
	"errors"
	"strconv"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

type floatCastNode struct {
	graph *g.Graph
	data  *g.NodeData
}

func BuildFloatCastNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	node := floatCastNode{
		graph: graph,
		data:  data,
	}

	return &node, nil
}

func (n *floatCastNode) Run(msg g.Msg) error {
	newMsg := make(g.Msg)

	payload, ok := msg["payload"].(string)
	if !ok {
		return errors.New("payload is not a string")
	}

	value, err := strconv.ParseFloat(payload, 32)
	if err != nil {
		return errors.New("can't parse value")
	}

	newMsg["payload"] = value
	return runChildren(msg, n.graph, n)
}

func (n *floatCastNode) Data() g.NodeData {
	return *n.data
}

func (n *floatCastNode) Type() g.NodeType {
	return g.FloatCastNode
}
