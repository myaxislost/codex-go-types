package protocol

type RpcMsg struct {
	Method Method `json:"method,omitempty"`
	Id     int64  `json:"id"`
	Params any    `json:"params,omitempty"`
	Result any    `json:"result,omitempty"`
	Error  any    `json:"error,omitempty"`
}

type InitializeParams struct {
	ClientInfo   ClientInfo              `json:"clientInfo"`
	Capabilities *InitializeCapabilities `json:"capabilities,omitempty"`
}

type InitializeResponse struct {
	UserAgent string `json:"userAgent"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Version string `json:"version"`
}

type InitializeCapabilities struct {
	ExperimentalApi           bool     `json:"experimentalApi,omitempty"`
	OptOutNotificationMethods []string `json:"optOutNotificationMethods,omitempty"`
}

type CollaborationModeType = string

const (
	PlanMode    CollaborationModeType = "plan"
	DefaultMode CollaborationModeType = "default"
)

type CollaborationMode struct {
	Mode     CollaborationModeType `json:"mode"`
	Settings TurnStartSettings     `json:"settings"`
	// ReasoningEffort *string `json:"reasoning_effort"`
}
type TurnStartSettings struct {
	DevInstructions *string `json:"developer_instructions"`
	Model           string  `json:"model"`
}

type Personality = string

const (
	PersonalityFriendly  Personality = "friendly"
	PersonalityPragmatic Personality = "pragmatic"
	PersonalityNone      Personality = "none"
)
