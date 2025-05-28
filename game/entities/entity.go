package entities

import (
	"block-rogue/game/uuidfactory"
	"math"
)

type Entity struct {
	UUID             string  `json:"uuid"`
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	DX               float64 `json:"dx"`
	DY               float64 `json:"dy"`
	Color            string  `json:"color"`
	Width            float64 `json:"width"`
	Height           float64 `json:"height"`
	Health           int     `json:"health"`
	Speed            float64 `json:"speed"`
	TraveledDistance float64 `json:"traveled_distance"`
	MaxDistance      float64 `json:"max_distance"`
	Type             string  `json:"type"`
	TimeAlive        int     `json:"time_alive"`
	MaxTimeAlive     int     `json:"max_time_alive"`
}

func (e *Entity) Move() {
	e.X += e.Speed * e.DX
	e.Y += e.Speed * e.DY
	e.TraveledDistance += e.Speed
	e.TimeAlive += 1
}

func NewProjectile(id int, name string, x, y, dx, dy float64, color string) *Entity {
	return &Entity{
		UUID:         uuidfactory.New(),
		ID:           id,
		Name:         name,
		X:            x,
		Y:            y,
		DX:           dx,
		DY:           dy,
		Color:        "black",
		Width:        5,
		Height:       5,
		Health:       1,   // Projectiles typically don't have health
		Speed:        5.0, // Speed of the projectile
		Type:         "projectile",
		TimeAlive:    0,      // Time alive in ticks
		MaxTimeAlive: 100000, // Example: alive for 100 ticks
	}
}

func NewEnemy(id int, name string, x, y float64) *Entity {
	return &Entity{
		UUID:      uuidfactory.New(),
		ID:        id,
		Name:      name,
		X:         x,
		Y:         y,
		DX:        0,
		DY:        0,
		Color:     "red",
		Width:     31,
		Height:    31,
		Health:    100,
		Speed:     2.0, // Speed of the enemy
		Type:      "enemy",
		TimeAlive: 0, // Time alive in ticks
	}
}

func (e *Entity) IsAlive() bool {
	if e.Health <= 0 {
		return false
	}
	if e.TraveledDistance >= e.MaxDistance {
		return false
	}
	if e.TimeAlive >= e.MaxTimeAlive {
		return false
	}
	if e.X < 0 || e.Y < 0 || e.X > 1000 || e.Y > 800 {
		return false
	}
	return true
}

func (e *Entity) FindPlayer(players map[string]*Player) {
	if e.Type != "enemy" {
		return
	}

	minDistance := math.MaxFloat64
	dx := 0.0
	dy := 0.0

	for _, player := range players {
		new_dx := player.X - e.X
		new_dy := player.Y - e.Y
		distance := new_dx*new_dx + new_dy*new_dy

		if distance < minDistance {
			minDistance = distance
			dx = new_dx
			dy = new_dy
		}
	}
	if minDistance == math.MaxFloat64 {
		return
	}
	if minDistance == 0 {
		e.DX = 0
		e.DY = 0
		return
	}
	if minDistance > 200*200 { // If the distance is less than 10 pixels, don't move
		e.DX = 0
		e.DY = 0
		return
	}
	norm := 1.0 / math.Sqrt(minDistance)
	dx *= norm
	dy *= norm
	e.DX = dx
	e.DY = dy

}

func (e *Entity) EnemyDamageFromProjectiles(projectiles map[string]*Entity) {
	for _, projectile := range projectiles {
		w := e.Width + projectile.Width
		w = w * w / 4
		if Distance(e, projectile) < w {
			e.Health -= projectile.Health
			projectile.Health = 0
		}
	}
}

func Distance(a, b *Entity) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}
