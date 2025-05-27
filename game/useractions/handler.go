package useractions

import (
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"encoding/json"
	"fmt"
	"math"
)

func HandleUserAction(a Action) error {
	t := a.Type

	switch t {

	case "attack":
		attack := Attack{}
		if err := json.Unmarshal(a.Action, &attack); err != nil {
			return fmt.Errorf("error unmarshalling attack action: %w", err)
		}
		gamestate.ActionQueue <- attack

	case "move":
		move := Move{}
		if err := json.Unmarshal(a.Action, &move); err != nil {
			return fmt.Errorf("error unmarshalling move action: %w", err)
		}
		gamestate.ActionQueue <- move

	case "chat":
		fmt.Printf("Chat action received: %s\n", a.Action)

	default:
		return fmt.Errorf("invalid action type: %s", t)
	}
	return nil
}

func (a Attack) Apply(state *gamestate.State) {
	dx := a.TargetX - a.SourceX
	dy := a.TargetY - a.SourceY
	norm := dx*dx + dy*dy
	norm = 1.0 / math.Sqrt(norm)
	dx *= norm
	dy *= norm
	entity := entities.Entity{
		ID:               len(state.Enemies),
		Name:             "attack",
		X:                a.SourceX + norm*20, // 10 pixels away from source
		Y:                a.SourceY + norm*20, // 10 pixels away from source
		DX:               dx,                  // 5 pixels per tick
		DY:               dy,                  // 5 pixels per tick
		Color:            "black",
		Width:            5,
		Height:           5,
		Health:           25,  // Attacks have no health
		Speed:            3.0, // Speed of the attack
		Type:             "projectile",
		TimeAlive:        0,      // Time alive in ticks
		MaxTimeAlive:     100000, // Example: alive for 100 ticks
		TraveledDistance: 0,
		MaxDistance:      400,
	}

	state.Projectiles = append(state.Projectiles, &entity)
}

func (a Move) Apply(state *gamestate.State) {
	for _, player := range state.Players {
		if player.ID == a.PlayerId {
			player.DX = a.DX
			player.DY = a.DY
			return
		}
	}
}
