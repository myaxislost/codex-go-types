package thread_start_response

import "encoding/json"

type ThreadStartResponse struct {
	ApprovalPolicy  string          `json:"approvalPolicy"`
	Cwd             string          `json:"cwd"`
	Model           string          `json:"model"`
	ModelProvider   string          `json:"modelProvider"`
	ReasoningEffort *string         `json:"reasoningEffort,omitempty"`
	Sandbox         json.RawMessage `json:"sandbox"`
	Thread          json.RawMessage `json:"thread"`
}
