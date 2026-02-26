package turn_completed_notification

import "encoding/json"

type TurnCompletedNotification struct {
	ThreadID string `json:"threadId"`
	Turn     Turn   `json:"turn"`
}

type Turn struct {
	ID     string          `json:"id,omitempty"`
	Status string          `json:"status,omitempty"`
	Error  *TurnError      `json:"error,omitempty"`
	Raw    json.RawMessage `json:"-"`
}

type TurnError struct {
	Code           string          `json:"code,omitempty"`
	Message        string          `json:"message,omitempty"`
	CodexErrorInfo *CodexErrorInfo `json:"codexErrorInfo,omitempty"`
}

type CodexErrorInfo struct {
	Code                           *string                                       `json:"code,omitempty"`
	HttpConnectionFailed           *HttpConnectionFailedCodexErrorInfo           `json:"httpConnectionFailed,omitempty"`
	ResponseStreamConnectionFailed *ResponseStreamConnectionFailedCodexErrorInfo `json:"responseStreamConnectionFailed,omitempty"`
	ResponseStreamDisconnected     *ResponseStreamDisconnectedCodexErrorInfo     `json:"responseStreamDisconnected,omitempty"`
	ResponseTooManyFailedAttempts  *ResponseTooManyFailedAttemptsCodexErrorInfo  `json:"responseTooManyFailedAttempts,omitempty"`
}

type HttpConnectionFailedCodexErrorInfo struct {
	HTTPStatusCode *uint16 `json:"httpStatusCode,omitempty"`
}

type ResponseStreamConnectionFailedCodexErrorInfo struct {
	HTTPStatusCode *uint16 `json:"httpStatusCode,omitempty"`
}

type ResponseStreamDisconnectedCodexErrorInfo struct {
	HTTPStatusCode *uint16 `json:"httpStatusCode,omitempty"`
}

type ResponseTooManyFailedAttemptsCodexErrorInfo struct {
	HTTPStatusCode *uint16 `json:"httpStatusCode,omitempty"`
}
