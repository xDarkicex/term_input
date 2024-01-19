package quiz

type Quiz struct {
	ID        int
	Title     string
	Questions []Question
}

type Question struct {
	Text       string
	Options    []string
	CorrectIdx int
}

func NewQuiz(title string) *Quiz {
	return &Quiz{
		Title:     title,
		Questions: []Question{},
	}
}

func NewQuestion(field string, options []string, correct int) *Question {
	return &Question{
		Text:       field,
		Options:    options,
		CorrectIdx: correct,
	}
}

func (q *Quiz) AddQuestion(question *Question) {
	q.Questions = append(q.Questions, *question)
}

func (q *Quiz) GetQuestions() []Question {
	return q.Questions
}
