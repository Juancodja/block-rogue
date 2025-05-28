package rutines

import (
	"block-rogue/game/entities"
	"block-rogue/game/gamestate"
	"block-rogue/game/uuidfactory"
	"math/rand"
	"time"
)

type CreateEnemy struct{}

func (a CreateEnemy) Apply(state *gamestate.State) {
	entity := entities.Entity{
		UUID:             uuidfactory.New(),
		ID:               len(state.Enemies),
		Name:             "enemy",
		X:                float64(rand.Intn(1000 - 10)), // hasta 990
		Y:                float64(rand.Intn(800 - 10)),  // hasta 790
		Color:            "red",
		Width:            64,
		Height:           64,
		DX:               0,
		DY:               0,
		Health:           100,
		Speed:            1.0, // px/tick
		Type:             "enemy",
		TimeAlive:        0,    // in ticks
		MaxTimeAlive:     5000, // ticks
		TraveledDistance: 0,
		MaxDistance:      100000, // px
	}
	state.Enemies[entity.UUID] = &entity
}

func StartEnemySpawner(state *gamestate.State) {
	ticker := time.NewTicker(200 * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		if len(state.Enemies) >= 50 {
			continue
		}
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
