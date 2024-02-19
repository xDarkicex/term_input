package player

import (
	"time"

	"github.com/xDarkicex/Development/term_input/quiz"
	"github.com/xDarkicex/Development/term_input/score"
)

type Player struct {
	ID                int
	Username          string
	Password          string
	CurrentQuiz       quiz.Quiz
	PreviousAttempts  []QuizAttempt
	CompletedQuizzes  []quiz.Quiz
	FailedQuizzes     []quiz.Quiz
	NextLevelProgress int
	AverageScore      float64
	ScoreboardRank    int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type QuizAttempt struct {
	QuizID    int
	Score     score.Score
	Completed bool
	Attempted time.Time
}

func NewPlayer() *Player {
	return &Player{
		ID:                0,
		Username:          "",
		Password:          "",
		CurrentQuiz:       quiz.Quiz{},
		PreviousAttempts:  []QuizAttempt{},
		CompletedQuizzes:  []quiz.Quiz{},
		FailedQuizzes:     []quiz.Quiz{},
		NextLevelProgress: 0,
		AverageScore:      0.0,
		ScoreboardRank:    0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func (p *Player) SetUsername(username string) {
	p.Username = username
}

func (p *Player) SetPassword(password string) {
	p.Password = password
}

func (p *Player) SetCurrentQuiz(quiz quiz.Quiz) {
	p.CurrentQuiz = quiz
}

func (p *Player) SetPreviousAttempts(attempt QuizAttempt) {
	p.PreviousAttempts = append(p.PreviousAttempts, attempt)
}

func (p *Player) SetCompletedQuizzes(quiz quiz.Quiz) {
	p.CompletedQuizzes = append(p.CompletedQuizzes, quiz)
}

func (p *Player) SetFailedQuizzes(quiz quiz.Quiz) {
	p.FailedQuizzes = append(p.FailedQuizzes, quiz)
}

func (p *Player) SetNextLevelProgress(progress int) {
	p.NextLevelProgress = p.NextLevelProgress + progress
}

func (p *Player) SetAverageScore(score float64) {
	p.AverageScore = p.AverageScore + score/float64(len(p.PreviousAttempts))
}

func (p *Player) SetScoreboardRank(rank int) {
	p.ScoreboardRank = rank
}

func (p *Player) GetUsername() string {
	return p.Username
}
func (p *Player) GetID() int {
	return p.ID
}

func (p *Player) GetPassword() string {
	return p.Password
}

func (p *Player) GetCurrentQuiz() quiz.Quiz {
	return p.CurrentQuiz
}

func (p *Player) GetPreviousAttempts() []QuizAttempt {
	return p.PreviousAttempts
}

func (p *Player) GetCompletedQuizzes() []quiz.Quiz {
	return p.CompletedQuizzes
}

func (p *Player) GetFailedQuizzes() []quiz.Quiz {
	return p.FailedQuizzes
}

func (p *Player) GetNextLevelProgress() int {
	return p.NextLevelProgress
}

func (p *Player) GetAverageScore() float64 {
	return p.AverageScore
}

func (p *Player) GetScoreboardRank() int {
	return p.ScoreboardRank
}

func (p *Player) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Player) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *Player) UpdateUsername(username string) {
	p.Username = username
}

func (p *Player) UpdatePassword(password string) {
	p.Password = password
}

func (p *Player) UpdateCurrentQuiz(quiz quiz.Quiz) {
	p.CurrentQuiz = quiz
}

func (p *Player) UpdatePreviousAttempts(attempt QuizAttempt) {
	p.PreviousAttempts = append(p.PreviousAttempts, attempt)
}

func (p *Player) UpdateCompletedQuizzes(quiz quiz.Quiz) {
	p.CompletedQuizzes = append(p.CompletedQuizzes, quiz)
}

func (p *Player) UpdateFailedQuizzes(quiz quiz.Quiz) {
	p.FailedQuizzes = append(p.FailedQuizzes, quiz)
}

func (p *Player) UpdateNextLevelProgress(progress int) {
	p.NextLevelProgress = progress
}

func (p *Player) UpdateAverageScore(score float64) {
	p.AverageScore = score
}

func (p *Player) UpdateScoreboardRank(rank int) {
	p.ScoreboardRank = rank
}

// Add any additional methods or functions related to player package here
