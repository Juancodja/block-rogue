package useractions

import (
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"block-rogue/game/uuidfactory"
	"math"
)

type Attack struct {
	PlayerId           int     `json:"player_id"`
	TargetX            float64 `json:"target_x"`
	TargetY            float64 `json:"target_y"`
	SourceX            float64 `json:"source_x"`
	SourceY            float64 `json:"source_y"`
	Width              float64 `json:"width"`
	Height             float64 `json:"height"`
	Damage             int     `json:"damage"`
	DistanceFromSource float64 `json:"distance_from_source"`
	TimeAlive          int     `json:"time_alive"`
	Speed              float64 `json:"speed"`
	MaxDistance        float64 `json:"max_distance"`
}

func (a Attack) Apply(state *gamestate.State) {
	dx := a.TargetX - a.SourceX
	dy := a.TargetY - a.SourceY
	norm := dx*dx + dy*dy
	norm = 1.0 / math.Sqrt(norm)
	dx *= norm
	dy *= norm
	entity := entities.Entity{
		UUID:             uuidfactory.New(),
		ID:               len(state.Enemies),
		Name:             "attack",
		X:                a.SourceX + dx*a.DistanceFromSource,
		Y:                a.SourceY + dy*a.DistanceFromSource,
		DX:               dx,
		DY:               dy,
		Color:            "black",
		Width:            a.Width,
		Height:           a.Height,
		Health:           a.Damage, // Attacks have no health
		Speed:            a.Speed,  // Speed of the attack
		Type:             "projectile",
		TimeAlive:        0,           // Time alive in ticks
		MaxTimeAlive:     a.TimeAlive, // Example: alive for 100 ticks
		TraveledDistance: 0,
		MaxDistance:      a.MaxDistance,
	}

	state.Projectiles[entity.UUID] = &entity
}
