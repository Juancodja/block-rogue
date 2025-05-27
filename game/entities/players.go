package entities

type Player struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	DX        float64 `json:"dx"`
	DY        float64 `json:"dy"`
	Color     string  `json:"color"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	MaxHealth int     `json:"max_health"`
	Health    int     `json:"health"`
	Speed     float64 `json:"speed"`
}

func (p *Player) Move() {
	p.X += p.Speed * p.DX
	p.Y += p.Speed * p.DY
}
