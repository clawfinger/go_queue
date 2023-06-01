package goqueue

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		arr:      make([]T, size+1),
		writeIdx: 0,
		readIdx:  0,
		size:     size + 1,
	}
}

type Queue[T any] struct {
	arr      []T
	size     int
	writeIdx int
	readIdx  int
}

func (q *Queue[T]) Enqueue(entry T) bool {
	if (q.writeIdx+1)%q.size == q.readIdx {
		return false
	}
	q.arr[q.writeIdx] = entry
	q.writeIdx = (q.writeIdx + 1) % q.size
	return true
}

func (q *Queue[T]) Deueue() (T, bool) {
	if q.writeIdx == q.readIdx {
		var res T
		return res, false
	}
	value := q.arr[q.readIdx]
	q.readIdx = (q.readIdx + 1) % q.size
	return value, true
}

func (q *Queue[T]) Empty() bool {
	return q.writeIdx == q.readIdx
}
