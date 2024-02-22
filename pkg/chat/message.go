package chat

type SocketMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}
