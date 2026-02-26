package config_read_response

import "encoding/json"

type ConfigReadResponse struct {
	Config  json.RawMessage            `json:"config"`
	Layers  []json.RawMessage          `json:"layers,omitempty"`
	Origins map[string]json.RawMessage `json:"origins"`
}
