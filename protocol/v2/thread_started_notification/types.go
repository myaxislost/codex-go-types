package thread_started_notification

import "encoding/json"

type ThreadStartedNotification struct {
	Thread json.RawMessage `json:"thread"`
}
