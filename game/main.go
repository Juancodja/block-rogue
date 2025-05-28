package main

import (
	"block-rogue/game/comms"
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"block-rogue/game/rutines"
	"block-rogue/game/uuidfactory"
	"time"
)

const VIEW_WIDTH = 1280
const VIEW_HEIGHT = 800

func main() {

	state := gamestate.State{
		Enemies: map[string]*entities.Entity{},
		Players: map[string]*entities.Player{},
	}

	player := entities.Player{
		UUID:      uuidfactory.New(),
		ID:        0,
		Name:      "Player1",
		X:         500 - 15,
		Y:         400 - 15,
		Color:     "blue",
		Width:     64,
		Height:    64,
		MaxHealth: 100,
		Health:    100,
		Speed:     1.0,
	}

	state.Players[player.UUID] = &player

	go comms.StartWebSocketServer("localhost", 8888)

	go rutines.StartEnemySpawner(&state)

	go rutines.StartEnemyAttack(&state)

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		comms.BroadcastState(&state)
		gamestate.UpdateState(&state)
	}
}
