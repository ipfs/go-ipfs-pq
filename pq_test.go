package pq

import (
	"sort"
	"testing"
)

type TestElem struct {
	Key      string
	Priority int
	index    int
}

func (e *TestElem) Index() int {
	return e.index
}

func (e *TestElem) SetIndex(i int) {
	e.index = i
}

var PriorityComparator = func(i, j Elem) bool {
	return i.(*TestElem).Priority > j.(*TestElem).Priority
}

func TestQueuesReturnTypeIsSameAsParameterToPush(t *testing.T) {
	q := New(PriorityComparator)
	expectedKey := "foo"
	elem := &TestElem{Key: expectedKey}
	q.Push(elem)
	switch v := q.Pop().(type) {
	case *TestElem:
		if v.Key != expectedKey {
			t.Fatal("the key doesn't match the pushed value")
		}
	default:
		t.Fatal("the queue is not casting values appropriately")
	}
}

func TestCorrectnessOfPop(t *testing.T) {
	q := New(PriorityComparator)
	tasks := []TestElem{
		{Key: "a", Priority: 9},
		{Key: "b", Priority: 4},
		{Key: "c", Priority: 3},
		{Key: "d", Priority: 0},
		{Key: "e", Priority: 6},
	}
	for _, e := range tasks {
		q.Push(&e)
	}
	var priorities []int
	var peekPriorities []int
	for q.Len() > 0 {
		i := q.Peek().(*TestElem).Priority
		peekPriorities = append(peekPriorities, i)
		j := q.Pop().(*TestElem).Priority
		t.Logf("popped %v", j)
		priorities = append(priorities, j)
	}
	if !sort.IntsAreSorted(peekPriorities) {
		t.Fatal("the values were not returned in sorted order")
	}
	if !sort.IntsAreSorted(priorities) {
		t.Fatal("the popped values were not returned in sorted order")
	}
}

func TestRemove(t *testing.T) {
	q := New(PriorityComparator)
	tasks := []TestElem{
		{Key: "a", Priority: 9},
		{Key: "b", Priority: 4},
		{Key: "c", Priority: 3},
	}
	for i := range tasks {
		q.Push(&tasks[i])
	}
	removed := q.Remove(1).(*TestElem)
	if q.Len() != 2 {
		t.Fatal("expected item to have been removed")
	}
	if q.Pop().(*TestElem).Key == removed.Key {
		t.Fatal("Remove() returned wrong element")
	}
	if q.Pop().(*TestElem).Key == removed.Key {
		t.Fatal("Remove() returned wrong element")
	}
}

func TestUpdate(t *testing.T) {
	t.Log(`
	Add 3 elements.
	Update the highest priority element to have the lowest priority and fix the queue.
	It should come out last.`)
	q := New(PriorityComparator)
	lowest := &TestElem{Key: "originallyLowest", Priority: 1}
	middle := &TestElem{Key: "originallyMiddle", Priority: 2}
	highest := &TestElem{Key: "toBeUpdated", Priority: 3}
	q.Push(middle)
	q.Push(highest)
	q.Push(lowest)
	if q.Peek().(*TestElem).Key != highest.Key {
		t.Fatal("head element doesn't have the highest priority")
	}
	if q.Pop().(*TestElem).Key != highest.Key {
		t.Fatal("popped element doesn't have the highest priority")
	}
	q.Push(highest)           // re-add the popped element
	highest.Priority = 0      // update the PQ
	q.Update(highest.Index()) // fix the PQ
	if q.Peek().(*TestElem).Key != middle.Key {
		t.Fatal("middle element should now have the highest priority")
	}
	if q.Pop().(*TestElem).Key != middle.Key {
		t.Fatal("middle element should now have the highest priority")
	}
}
