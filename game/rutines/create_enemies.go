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
		ID:     len(state.Entities),
		Name:   "enemy",
		X:      float64(rand.Intn(1000 - 10)), // hasta 990
		Y:      float64(rand.Intn(800 - 10)),  // hasta 790
		Color:  "red",
		Width:  20,
		Height: 20,
	}
	state.Entities = append(state.Entities, entity)
}

func StartEnemySpawner(state *gamestate.State) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		action := CreateEnemy{}
		gamestate.ActionQueue <- action
	}
}
