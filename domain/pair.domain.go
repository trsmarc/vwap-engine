package domain

import "fmt"

type Pair struct {
	From string
	To   string
}

func NewPair(from string, to string) Pair {
	return Pair{From: from, To: to}
}

func (t Pair) String() string {
	return fmt.Sprintf("%s-%s", t.From, t.To)
}
