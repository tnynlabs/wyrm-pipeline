package nodes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	g "github.com/tnynlabs/wyrm-pipeline/pkg/graph"
)

const INVOKE_URL = "http://%s/api/v1/devices/%d/invoke/%s"

type deviceInvokeNode struct {
	graph    *g.Graph
	data     *g.NodeData
	deviceID int64
	pattern  string
}

func BuildDeviceInvokeNode(graph *g.Graph, data *g.NodeData) (g.Node, error) {
	deviceID, ok := data.Attrs["device_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("'device_id' attribute is required")
	}

	pattern, ok := data.Attrs["pattern"].(string)
	if !ok {
		return nil, fmt.Errorf("'pattern' attribute is required")
	}

	node := deviceInvokeNode{
		graph:    graph,
		data:     data,
		deviceID: int64(deviceID),
		pattern:  pattern,
	}

	return &node, nil
}

func (n *deviceInvokeNode) Run(msg g.Msg) error {
	payload, err := json.Marshal(msg["payload"])
	if err != nil {
		return err
	}
	body := strings.NewReader(string(payload))

	wyrmApiHost := os.Getenv("WYRM_API_HOST")
	wyrmApiPort := os.Getenv("WYRM_API_PORT")
	wyrmAddr := wyrmApiHost + ":" + wyrmApiPort
	url := fmt.Sprintf(INVOKE_URL, wyrmAddr, n.deviceID, n.pattern)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	newMsg := make(g.Msg)
	newMsg["payload"] = string(respBody)

	return runChildren(newMsg, n.graph, n)
}

func (n *deviceInvokeNode) Data() g.NodeData {
	return *n.data
}

func (n *deviceInvokeNode) Type() g.NodeType {
	return g.DeviceInvokeNode
}
