package useractions

import "encoding/json"

type Action struct {
	UserId int             `json:"user_id"`
	Type   string          `json:"type"`
	Action json.RawMessage `json:"action"`
}

type Attack struct {
	PlayerId int     `json:"player_id"`
	TargetX  float64 `json:"target_x"`
	TargetY  float64 `json:"target_y"`
	SourceX  float64 `json:"source_x"`
	SourceY  float64 `json:"source_y"`
}

type Move struct {
	PlayerId int     `json:"player_id"`
	DX       float64 `json:"dx"`
	DY       float64 `json:"dy"`
}
