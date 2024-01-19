package updater

import (
	"github.com/xDarkicex/Development/term_input/score"
	"github.com/xDarkicex/Development/term_input/scoreboard"
)

type Updater struct {
	Scoreboard *scoreboard.Scoreboard
	Score      *score.Score
}

func NewUpdater() *Updater {
	return &Updater{
		Scoreboard: scoreboard.NewScoreboard(),
		Score:      score.NewScore(),
	}
}
