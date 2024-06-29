package utilities

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() *T {
	if len(*q) == 0 {
		return nil
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return &v
}

func (q *Queue[T]) Peek() *T {
	if len(*q) == 0 {
		return nil
	}
	return &(*q)[0]
}

func (q *Queue[T]) PeekAll() *[]T {
	return (*[]T)(q)
}
