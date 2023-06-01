package goqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInterface interface {
	Get() int
}

type TestStruct struct {
	val int
}

func (t *TestStruct) Get() int {
	return t.val
}

func New(val int) *TestStruct {
	return &TestStruct{
		val: val,
	}
}
func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[TestInterface](3)
	res := q.Enqueue(New(1))
	assert.True(t, res)
	res = q.Enqueue(New(2))
	assert.True(t, res)
	res = q.Enqueue(New(3))
	assert.True(t, res)
	res = q.Enqueue(New(4))
	assert.False(t, res)
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[TestInterface](3)
	res := q.Enqueue(New(1))
	assert.True(t, res)
	res = q.Enqueue(New(2))
	assert.True(t, res)
	ret, ok := q.Deueue()
	assert.True(t, ok)
	assert.Equal(t, 1, ret.Get())
	ret, ok = q.Deueue()
	assert.True(t, ok)
	assert.Equal(t, 2, ret.Get())
	_, ok = q.Deueue()
	assert.False(t, ok)
}
