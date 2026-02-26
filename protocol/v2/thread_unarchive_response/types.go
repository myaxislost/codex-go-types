package thread_unarchive_response

import "encoding/json"

type ThreadUnarchiveResponse struct {
	Thread json.RawMessage `json:"thread"`
}
