package graph

type NodeType string
type Msg map[string]interface{}

const (
	// Trigger Nodes
	WebhookTriggerNode NodeType = "WEBHOOK_TRIGGER"

	FloatCastNode    NodeType = "FLOAT_CAST"
	EmailNode        NodeType = "EMAIL"
	ConditionNode    NodeType = "CONDITION"
	DeviceInvokeNode NodeType = "INVOKE_DEVICE"
)

type Node interface {
	Run(msg Msg) error
	Data() NodeData
	Type() NodeType
}

type NodeData struct {
	Id       string                 `json:"id"`
	Type     NodeType               `json:"type"`
	Attrs    map[string]interface{} `json:"attributes"`
	Children []string               `json:"children"`
}
