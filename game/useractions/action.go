package useractions

import "encoding/json"

type Action struct {
	UserId int             `json:"user_id"`
	Type   string          `json:"type"`
	Action json.RawMessage `json:"action"`
}
