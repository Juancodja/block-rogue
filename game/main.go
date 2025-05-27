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
		Entities: []*entities.Entity{
			{
				ID:     0,
				Name:   "tree1",
				X:      100,
				Y:      100,
				DX:     0,
				DY:     0,
				Width:  50,
				Height: 50,
				Color:  "green",
			},
		},
		Players: []*entities.Player{
			{
				ID:        1,
				Name:      "Player1",
				X:         500 - 15,
				Y:         400 - 15,
				Color:     "red",
				Width:     30,
				Height:    30,
				MaxHealth: 100,
				Health:    100,
				Speed:     5.0,
			},
		},
	}

	go comms.StartWebSocketServer("localhost", 8888)

	go rutines.StartEnemySpawner(&state)

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		comms.BroadcastState(&state)
		gamestate.UpdateState(&state)
	}
}
