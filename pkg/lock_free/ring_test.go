package lockfree

import (
	"reflect"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](10)

	for i := 0; i < 20; i++ {
		ok := rb.Put(i)
		if i < 10 && !ok {
			t.Errorf("put failed, %d:%t", i, ok)
		}
		if i > 9 && ok {
			t.Errorf("put failed, %d:%t", i, ok)
		}
	}
	v := rb.LookAll()
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(v, want) {
		t.Errorf("not equal: %v", v)
	}

	for i := 0; i < 5; i++ {
		v := rb.Get()
		if v != i {
			t.Errorf("get failed, %d:%d", v, i)
		}
	}

	v = rb.LookAll()
	want = []int{5, 6, 7, 8, 9}
	if !reflect.DeepEqual(v, want) {
		t.Errorf("not equal")
	}
}
