package thread_resume_response

import "encoding/json"

type ThreadResumeResponse struct {
	ApprovalPolicy  string          `json:"approvalPolicy"`
	Cwd             string          `json:"cwd"`
	Model           string          `json:"model"`
	ModelProvider   string          `json:"modelProvider"`
	ReasoningEffort *string         `json:"reasoningEffort,omitempty"`
	Sandbox         json.RawMessage `json:"sandbox"`
	Thread          json.RawMessage `json:"thread"`
}
