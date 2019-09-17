package queue

import (
	"testing"
)

type sitem struct {
	i int
}

func TestStackNew(t *testing.T) {
	s := Stack{}
	if s.Len() != 0 {
		t.Error("empty stack shold have length of zero")
	}
	if s.Pop() != nil {
		t.Error("empty stack should pop out nils #1")
	}
	if s.Pop() != nil {
		t.Error("empty stack should pop out nils #2")
	}
}

func TestStackPushPop(t *testing.T) {
	a := &sitem{}
	b := &sitem{}
	c := &sitem{}

	s := Stack{}

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

	if s.Pop() != c {
		t.Error("pop C node failed")
	}
	if s.Pop() != b {
		t.Error("pop B node failed")
	}
	if s.Pop() != a {
		t.Error("pop A node failed")
	}
	if s.Pop() != nil {
		t.Error("pop from empty stack failed")
	}
}

func TestStackPushNil(t *testing.T) {
	s := Stack{}
	a := &sitem{}
	b := &sitem{}
	s.Push(a)
	s.Push(nil)
	s.Push(b)
	if s.Len() != 2 {
		t.Error("Push() should ignore nils")
	}

	if b != s.Pop() || a != s.Pop() {
		t.Error("a & b not returned")
	}
}

func TestStackConcurency(t *testing.T) {
	s := Stack{}
	count := 100
	nodes := make([]*sitem, count)
	for i := 0; i < count; i++ {
		nodes[i] = &sitem{}
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

func BenchmarkStackPush(b *testing.B) {
	b.StopTimer()
	s := Stack{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := &sitem{i: i}
		s.Push(n)
	}
}

func BenchmarkStackPop(b *testing.B) {
	b.StopTimer()
	s := Stack{}
	for i := 0; i < b.N; i++ {
		n := &sitem{i: i}
		s.Push(n)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
