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
