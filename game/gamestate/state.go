package gamestate

import e "block-rogue/game/entities"

type State struct {
	Entities    []*e.Entity `json:"entities"`
	Players     []*e.Player `json:"players"`
	Proyectiles []*e.Entity `json:"projectiles"`
}
