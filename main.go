package main

import (
	"fmt"
	"log"

	"github.com/xDarkicex/Development/term_input/capture"
	"github.com/xDarkicex/Development/term_input/quiz"
	"github.com/xDarkicex/Development/term_input/updater"
)

func init() {

}

func main() {
	update := updater.NewUpdater()
	q := quiz.NewQuiz("Math Quiz")
	input := capture.NewInput()
	q.AddQuestion(quiz.NewQuestion("What is 1+1?", []string{"2", "3", "1", "6"}, 1))
	q.AddQuestion(quiz.NewQuestion("What is 2+2?", []string{"12", "2", "4", "6"}, 3))
	q.AddQuestion(quiz.NewQuestion("What is 3+3?", []string{"10", "12", "5", "6"}, 4))
	q.AddQuestion(quiz.NewQuestion("What is 4+4?", []string{"9", "8", "13", "14"}, 2))
	q.AddQuestion(quiz.NewQuestion("What is 5+5?", []string{"7", "8", "10", "14"}, 3))

	fmt.Printf("%s\n", q.Title)

	for _, question := range q.GetQuestions() {
		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Println(i+1, ":", option)
		}

		fmt.Print("Enter Selection: ")
		selection, err := input.ReadInt()
		if err != nil {
			log.Println(err)
			continue
		}

		if selection < 1 || selection > len(question.Options) {
			fmt.Println("Invalid selection!")
			continue
		}

		if selection == question.CorrectIdx {
			fmt.Println("Correct!")
			update.Score.Add()
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("Your score is: %.1f out of: %d\n", update.Score.Value, len(q.GetQuestions()))
}
