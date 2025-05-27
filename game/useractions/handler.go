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
	case "move":

	case "attack":
		attack := Attack{}
		if err := json.Unmarshal(a.Action, &attack); err != nil {
			return fmt.Errorf("error unmarshalling attack action: %w", err)
		}
		gamestate.ActionQueue <- attack
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

	entity := entities.Entity{
		ID:     len(state.Entities),
		Name:   "attack",
		X:      a.SourceX + norm*20, // 10 pixels away from source
		Y:      a.SourceY + norm*20, // 10 pixels away from source
		DX:     dx * norm * 5,       // 5 pixels per tick
		DY:     dy * norm * 5,       // 5 pixels per tick
		Color:  "black",
		Width:  5,
		Height: 5,
	}

	state.Entities = append(state.Entities, &entity)
}
