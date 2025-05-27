package rutines

import (
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"math/rand"
	"time"
)

type CreateEnemy struct{}

func (a CreateEnemy) Apply(state *gamestate.State) {
	entity := &entities.Entity{
		ID:               len(state.Enemies),
		Name:             "enemy",
		X:                float64(rand.Intn(1000 - 10)), // hasta 990
		Y:                float64(rand.Intn(800 - 10)),  // hasta 790
		Color:            "red",
		Width:            20,
		Height:           20,
		DX:               0,
		DY:               0,
		Health:           100,
		Speed:            1.0, // Speed of the enemy
		Type:             "enemy",
		TimeAlive:        0,    // Time alive in ticks
		MaxTimeAlive:     5000, // Example: alive for 100 ticks
		TraveledDistance: 0,
		MaxDistance:      100000, // Example: max distance traveled
	}
	state.Enemies = append(state.Enemies, entity)
}

func StartEnemySpawner(state *gamestate.State) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		action := CreateEnemy{}
		gamestate.ActionQueue <- action
	}
}

type EnemyAttackPlayer struct{}

func (a EnemyAttackPlayer) Apply(state *gamestate.State) {
	for _, entity := range state.Enemies {
		entity.FindPlayer(state.Players)
	}
}

func StartEnemyAttack(state *gamestate.State) {
	ticker := time.NewTicker(400 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		action := EnemyAttackPlayer{}
		gamestate.ActionQueue <- action
	}
}
