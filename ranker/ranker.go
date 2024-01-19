package ranker

import (
	"github.com/xDarkicex/Development/term_input/player"
	"github.com/xDarkicex/Development/term_input/scoreboard"
)

func Rank(scoreboard *scoreboard.Scoreboard) []player.Player {
	players := scoreboard.Players
	for i := 0; i < len(players); i++ {
		for j := 0; j < len(players); j++ {
			if players[i].AverageScore > players[j].AverageScore {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
	return players
}
