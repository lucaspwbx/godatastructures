package queue

type Queue struct {
	dataStore []int
}

func (q *Queue) Enqueue(element int) error {
	q.dataStore = append(q.dataStore, element)
	return nil
}

func (q *Queue) Dequeue() error {
	q.dataStore = q.dataStore[:len(q.dataStore)-1]
	return nil
}

func (q *Queue) Front() int {
	return q.dataStore[0]
}

func (q *Queue) Back() int {
	return q.dataStore[len(q.dataStore)-1]
}

func (q *Queue) Empty() bool {
	if len(q.dataStore) == 0 {
		return true
	}
	return false
}

func NewQueue() *Queue {
	return &Queue{}
}
