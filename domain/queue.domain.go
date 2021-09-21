package domain

type Queue struct {
	tradings []Trading
}

func NewQueue() *Queue {
	return &Queue{tradings: make([]Trading, 0)}
}

func (t *Queue) Add(trading Trading) {
	t.tradings = append(t.tradings, trading)
}

func (t *Queue) Remove() Trading {
	var first = t.tradings[0]
	t.tradings = t.tradings[1:]
	return first
}

func (t *Queue) Len() int {
	return len(t.tradings)
}
