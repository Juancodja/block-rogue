package useractions

import (
	"block-rogue/game/gamestate"
	"encoding/json"
	"fmt"
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
