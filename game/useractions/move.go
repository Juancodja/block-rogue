package useractions

import "block-rogue/game/gamestate"

type Move struct {
	PlayerId int     `json:"player_id"`
	DX       float64 `json:"dx"`
	DY       float64 `json:"dy"`
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
