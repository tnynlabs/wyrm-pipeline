package nodes

import (
	"crypto/tls"
	"fmt"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
	"gopkg.in/gomail.v2"
)

type emailNode struct {
	graph   *g.Graph
	data    *g.NodeData
	to      string
	subject string
	body    string
}

func BuildEmailNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	to, ok := data.Attrs["to"].(string)
	if !ok {
		return nil, fmt.Errorf("'to' attribute is required")
	}
	subject, ok := data.Attrs["subject"].(string)
	if !ok {
		return nil, fmt.Errorf("'subject' attribute is required")
	}
	body, ok := data.Attrs["body"].(string)
	if !ok {
		return nil, fmt.Errorf("'body' attribute is required")
	}

	node := emailNode{
		graph:   graph,
		data:    data,
		to:      to,
		subject: subject,
		body:    body,
	}

	return &node, nil
}

func (n *emailNode) Run(msg g.Msg) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "wyrm.tnynlabs@gmail.com")
	m.SetHeader("To", n.to)
	m.SetHeader("Subject", n.subject)
	m.SetBody("text/html", n.body)

	d := gomail.NewDialer("smtp.gmail.com", 465, "wyrm.tnynlabs", "wosjelniifcvfhbd")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return runChildren(msg, n.graph, n)
}

func (n *emailNode) Data() g.NodeData {
	return *n.data
}

func (n *emailNode) Type() g.NodeType {
	return g.EmailNode
}
