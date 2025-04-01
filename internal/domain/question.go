package domain

type Question struct {
	ID       int
	Question string
	Answers  []string
	TrueAns  int
}
