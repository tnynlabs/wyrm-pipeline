package nodes

import (
	"log"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

type webhookTriggerNode struct {
	graph *g.Graph
	data  *g.NodeData
}

func BuildWebhookTriggerNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	node := webhookTriggerNode{
		graph: graph,
		data:  data,
	}

	return &node, nil
}

func (n *webhookTriggerNode) Run(msg g.Msg) error {
	log.Println(msg)

	return runChildren(msg, n.graph, n)
}

func (n *webhookTriggerNode) Data() g.NodeData {
	return *n.data
}

func (n *webhookTriggerNode) Type() g.NodeType {
	return g.WebhookTriggerNode
}
