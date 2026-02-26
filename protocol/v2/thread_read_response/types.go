package thread_read_response

import "encoding/json"

type ThreadReadResponse struct {
	Thread json.RawMessage `json:"thread"`
}
