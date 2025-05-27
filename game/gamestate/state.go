package gamestate

import e "block-rogue/game/entities"

type State struct {
	Enemies     []*e.Entity `json:"enemies"`
	Players     []*e.Player `json:"players"`
	Projectiles []*e.Entity `json:"projectiles"`
}
