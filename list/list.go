package list

import "errors"

type List struct {
	ds  []interface{}
	pos int
}

func (l *List) Insert(element interface{}) {
	l.ds = append(l.ds, element)
	l.pos++
}

func (l *List) FindPos(element interface{}) (int, error) {
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
	l.pos--
}

func (l *List) Clear() {
	l.ds = nil
}

func (l *List) Size() int {
	return len(l.ds)
}

func (l *List) Contains(element interface{}) bool {
	for _, v := range l.ds {
		if v == element {
			return true
		}
	}
	return false
}

func (l *List) Front() {
	l.pos = 0
}

func (l *List) End() {
	l.pos = l.Size() - 1
}

func (l *List) CurrentPos() int {
	return l.pos
}

func (l *List) Prev() {
	l.pos--
}

func (l *List) Next() {
	l.pos++
}

func (l *List) MoveTo(position int) {
	l.pos = position
}

func (l *List) GetElement() interface{} {
	return l.ds[l.pos]
}

func NewList() *List {
	return &List{pos: 0}
}
