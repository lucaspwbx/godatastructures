package list

import "errors"

type List struct {
	ds []int
}

func (l *List) Insert(element int) error {
	l.ds = append(l.ds, element)
	return nil
}

func (l *List) FindPos(element int) (int, error) {
	for k, v := range l.ds {
		if v == element {
			return k, nil
		}
	}
	return 0, errors.New("Element not found")
}

func (l *List) Remove(element int) {
	i, err := l.FindPos(element)
	if err != nil {
		return
	}
	l.ds = append(l.ds[:i], l.ds[i+1:]...)
}

func NewList() *List {
	return &List{}
}
