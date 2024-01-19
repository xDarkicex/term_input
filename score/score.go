package score

type Score struct {
	Value float64
}

func NewScore() *Score {
	return &Score{
		Value: 0,
	}
}

func (s *Score) Add() {
	s.Value = s.Value + 1
}
