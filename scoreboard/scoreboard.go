package scoreboard

import "github.com/xDarkicex/Development/term_input/player"

type Scoreboard struct {
	Players []player.Player
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{
		Players: []player.Player{},
	}
}
