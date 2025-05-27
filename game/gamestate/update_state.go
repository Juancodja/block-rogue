package gamestate

func UpdateState(state *State) {
	ConsumeActionQueue(state)
	UpdatePosition(state)
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

func UpdatePosition(state *State) {
	for _, entity := range state.Entities {
		entity.UpdatePosition()
	}
}
