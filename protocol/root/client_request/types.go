package client_request

import "encoding/json"

// ClientRequest is a generic JSON-RPC client request envelope.
type ClientRequest struct {
	ID     json.RawMessage `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}
