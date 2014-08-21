package list

import "errors"

type List struct {
	ds []interface{}
}

func (l *List) Insert(element interface{}) {
	l.ds = append(l.ds, element)
}

func (l *List) FindPos(element interface {}) (int, error) {
	for k, v := range l.ds {
		if v == element {
			return k, nil
		}
	}
	return 0, errors.New("Element not found")
}

func (l *List) Remove(element interface{}) {
	i, err := l.FindPos(element)
	if err != nil {
		return
	}
	l.ds = append(l.ds[:i], l.ds[i+1:]...)
}

func (l *List) Clear() {
	l.ds = nil
}

func (l *List) Size() int {
	return len(l.ds)
}

func NewList() *List {
	return &List{}
}
