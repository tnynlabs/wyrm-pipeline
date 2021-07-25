package nodes

import (
	"fmt"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/vm"
)

type conditionNode struct {
	graph *g.Graph
	data  *g.NodeData
	expr  expr.Node
}

// TODO: Node attributes validation

func BuildConditionNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	exprRaw, ok := data.Attrs["expr"].(string)
	if !ok {
		return nil, fmt.Errorf("'expr' attribute is required")
	}

	exprNode, err := expr.ParseExpression(exprRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid 'expr' attribute formatting")
	}

	node := conditionNode{
		graph: graph,
		data:  data,
		expr:  exprNode,
	}

	return &node, nil
}

func (n *conditionNode) Run(msg g.Msg) error {
	evalContext := datasource.NewContextSimpleNative(msg)
	val, ok := vm.Eval(evalContext, n.expr)
	if !ok {
		return fmt.Errorf("error evaluating expression")
	}

	if boolVal, ok := val.Value().(bool); ok && boolVal {
		return runChildren(msg, n.graph, n)
	}

	return nil
}

func (n *conditionNode) Data() g.NodeData {
	return *n.data
}

func (n *conditionNode) Type() g.NodeType {
	return g.ConditionNode
}
