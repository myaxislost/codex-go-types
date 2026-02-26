package thread_rollback_response

import "encoding/json"

type ThreadRollbackResponse struct {
	Thread json.RawMessage `json:"thread"`
}
