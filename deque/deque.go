package deque

type Deque struct {
	ds []interface{}
}

func (d *Deque) Size() int {
	return len(d.ds)
}

func (d *Deque) IsEmpty() bool {
	return d.Size() == 0
}

func (d *Deque) AddRear(element interface{}) {
	var newSlice []interface{}
	currSlice := d.ds
	newSlice = append(newSlice, element)
	newSlice = append(newSlice, currSlice...)
	d.ds = newSlice
}

func (d *Deque) AddFront(element interface{}) {
	d.ds = append(d.ds, element)
}

func (d *Deque) RemoveRear() {
	d.ds = d.ds[1:]
}

func (d *Deque) RemoveFront() {
	d.ds = d.ds[:d.Size()-1]
}

func NewDeque() *Deque {
	return &Deque{}
}
