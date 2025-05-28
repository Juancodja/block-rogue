package gamestate

import e "block-rogue/game/entities"

type State struct {
	Enemies     map[string]*e.Entity `json:"enemies"`
	Players     map[string]*e.Player `json:"players"`
	Projectiles map[string]*e.Entity `json:"projectiles"`
}
