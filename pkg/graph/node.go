package graph

type NodeType string

const(
	DeviceResponseNode	NodeType = "DEVICE_RESPONSE"
	ConditionNode		NodeType = "CONDITION"
	DeviceInvokeNode	NodeType = "DEVICE_INVOKE"
)

type Node interface {
	Run(input string) error
	Data() NodeData
	Type() NodeType
}

type NodeData struct {
	Id string						`json:"id"`
	Type NodeType					`json:"type"`
	Attrs map[string]interface{}	`json:"attributes"`
	Children []string				`json:"children"`
}
