package gamestate

type MutatesState interface {
	Apply(state *State)
}

var ActionQueue = make(chan MutatesState, 1000)
