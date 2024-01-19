package grader

import (
	"github.com/xDarkicex/Development/term_input/quiz"
)

func Grade(q quiz.Quiz, a []int) float64 {
	var score float64
	for i, question := range q.GetQuestions() {
		if question.CorrectIdx == a[i] {
			score++
		}
	}
	return score / float64(len(q.GetQuestions()))
}
