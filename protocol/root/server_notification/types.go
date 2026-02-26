package server_notification

import "encoding/json"

// ServerNotification is a generic JSON-RPC server notification envelope.
type ServerNotification struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}
