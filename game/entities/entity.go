package entities

type Entity struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	DX     float64 `json:"dx"`
	DY     float64 `json:"dy"`
	Color  string  `json:"color"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (e *Entity) UpdatePosition() {
	e.X += e.DX
	e.Y += e.DY
}
