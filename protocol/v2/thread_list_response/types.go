package thread_list_response

import "encoding/json"

type ThreadListResponse struct {
	Data       []json.RawMessage `json:"data"`
	NextCursor *string           `json:"nextCursor,omitempty"`
}
