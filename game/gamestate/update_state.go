package gamestate

import "block-rogue/game/entities"

func UpdateState(state *State) {
	ConsumeActionQueue(state)
	UpdateEnemies(state)
	UpdatePlayers(state)
	UpdateProjectiles(state)
	DamageEnemies(state)
	state.Enemies = FilterAlive(state.Enemies)
	state.Projectiles = FilterAlive(state.Projectiles)
}

func ConsumeActionQueue(state *State) {
	for {
		select {
		case mutate := <-ActionQueue:
			mutate.Apply(state)
		default:
			return
		}
	}
}

func DamageEnemies(state *State) {
	for _, e := range state.Enemies {
		e.EnemyDamageFromProjectiles(state.Projectiles)
	}
}

func UpdateEnemies(state *State) {
	for _, entity := range state.Enemies {
		entity.Move()
	}
}

func UpdatePlayers(state *State) {
	for _, player := range state.Players {
		player.Move()
	}
}

func UpdateProjectiles(state *State) {
	for _, projectile := range state.Projectiles {
		projectile.Move()
	}
}
func FilterAlive(entities []*entities.Entity) []*entities.Entity {
	n := 0
	for _, e := range entities {
		if e.IsAlive() {
			entities[n] = e
			n++
		}
	}
	for i := n; i < len(entities); i++ {
		entities[i] = nil
	}
	return entities[:n]
}
