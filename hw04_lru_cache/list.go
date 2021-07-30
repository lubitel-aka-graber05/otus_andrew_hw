package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	// Remove me after realization.
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front == nil {
		n := &ListItem{Value: v, Next: nil, Prev: nil}
		l.front = n
		l.back = n
	} else {
		n := &ListItem{Value: v, Next: l.front, Prev: nil}
		l.front.Prev = n
		l.front = n
	}
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.back == nil {
		n := &ListItem{Value: v, Next: nil, Prev: nil}
		l.front = n
		l.back = n
	} else {
		n := &ListItem{Value: v, Next: nil, Prev: l.back}
		l.back.Next = n
		l.back = n
	}
	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	switch i {
	case l.front:
		if l.front.Next != nil {
			l.front = l.front.Next
			l.front.Prev = nil
		} else {
			l.front, l.back = nil, nil
		}
	case l.back:
		if l.back.Prev != nil {
			l.back = l.back.Prev
			l.back.Next = nil
		}
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		// l.front = i
		return
	}
	if i == l.back {
		l.front = l.PushFront(i.Value)
		l.Remove(i)
	}
	if i == l.front.Next {
		l.front = l.PushFront(i.Value)
		l.Remove(i)
	}
}

func NewList() List {
	return new(list)
}
