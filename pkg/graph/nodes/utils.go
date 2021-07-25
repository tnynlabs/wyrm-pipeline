package nodes

import (
	"fmt"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

func runChildren(msg g.Msg, graph *g.Graph, node g.Node) error {
	for _, childId := range node.Data().Children {
		childNode, err := graph.GetNode(childId)
		if err != nil {
			return err
		}

		err = childNode.Run(msg)
		if err != nil {
			return fmt.Errorf("node %s: %w", childId, err)
		}
	}

	return nil
}
