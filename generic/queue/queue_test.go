package queue

import (
	"testing"
)

type qitem struct {
	i int
}

func TestQueueNew(t *testing.T) {
	s := Queue{}
	if s.Len() != 0 {
		t.Error("empty queue shold have length of zero")
	}
	if s.Pop() != nil {
		t.Error("empty queue should pop out nils #1")
	}
	if s.Pop() != nil {
		t.Error("empty queue should pop out nils #2")
	}
}

func TestQueuePushPop(t *testing.T) {
	a := &qitem{}
	b := &qitem{}
	c := &qitem{}

	s := Queue{}

	s.Push(a)
	if s.Len() != 1 {
		t.Error("push A failed")
	}

	s.Push(b)
	if s.Len() != 2 {
		t.Error("push B failed")
	}

	s.Push(c)
	if s.Len() != 3 {
		t.Error("push C failed")
	}

	if s.Pop() != a {
		t.Error("pop C node failed")
	}
	if s.Pop() != b {
		t.Error("pop B node failed")
	}
	if s.Pop() != c {
		t.Error("pop A node failed")
	}
	if s.Pop() != nil {
		t.Error("pop from empty йгугу failed")
	}
}

func TestQueuePushNil(t *testing.T) {
	s := Queue{}
	a := &qitem{}
	b := &qitem{}
	s.Push(a)
	s.Push(nil)
	s.Push(b)
	if s.Len() != 2 {
		t.Error("Push() should ignore nils")
	}

	if a != s.Pop() || b != s.Pop() {
		t.Error("a & b not returned")
	}
}

func TestQueueConcurency(t *testing.T) {
	s := Queue{}
	count := 100
	nodes := make([]*qitem, count)
	for i := 0; i < count; i++ {
		nodes[i] = &qitem{}
	}

	lenchan := make(chan int)

	for _, n := range nodes {
		n := n
		go func() {
			lenchan <- s.Push(n)
		}()
	}

	lens := make([]bool, 2*count)

	for i := 0; i < count; i++ {
		l := <-lenchan
		if lens[l] {
			t.Errorf("Push() should be thread safe; l=%d", l)
			return
		}
		lens[l] = true
	}
}

func BenchmarkQueuePush(b *testing.B) {
	b.StopTimer()
	s := Queue{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := &qitem{i: i}
		s.Push(n)
	}
}

func BenchmarkQueuePop(b *testing.B) {
	b.StopTimer()
	s := Queue{}
	for i := 0; i < b.N; i++ {
		n := &qitem{i: i}
		s.Push(n)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
