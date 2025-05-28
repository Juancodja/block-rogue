package main

import (
	"block-rogue/game/comms"
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"block-rogue/game/rutines"
	"time"
)

func main() {

	state := gamestate.State{
		Enemies: []*entities.Entity{},
		Players: []*entities.Player{
			{
				ID:        0,
				Name:      "Player1",
				X:         500 - 15,
				Y:         400 - 15,
				Color:     "blue",
				Width:     33,
				Height:    33,
				MaxHealth: 100,
				Health:    100,
				Speed:     1.0,
			},
		},
	}

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
