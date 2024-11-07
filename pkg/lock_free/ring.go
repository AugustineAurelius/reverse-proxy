package lockfree

import "sync"

type RingBuffer[T any] struct {
	buf        []T
	mu         sync.Mutex
	head, tail int
	len, cap   int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, capacity),
		mu:  sync.Mutex{},
		cap: capacity,
	}
}

// Put puts x into ring buffer
func (rb *RingBuffer[T]) Put(x T) (ok bool) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.len == rb.cap {
		return
	}

	rb.buf[rb.tail] = x
	rb.tail++
	if rb.tail > rb.cap-1 {
		rb.tail = 0
	}
	rb.len++
	ok = true
	return
}

// Get gets the first element from queue
func (rb *RingBuffer[T]) Get() (x T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	x = rb.buf[rb.head]
	rb.head++
	if rb.head > rb.cap-1 {
		rb.head = 0
	}
	rb.len--
	return
}

// IsFull checks if the ring buffer is full
func (rb *RingBuffer[T]) IsFull() bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	return rb.len == rb.cap
}

func (rb *RingBuffer[T]) LookAll() []T {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	all := make([]T, rb.len)
	j := 0
	for i := rb.head; ; i++ {
		if i > rb.cap-1 {
			i = 0
		}
		if i == rb.tail && j > 0 {
			break
		}
		all[j] = rb.buf[i]
		j++
	}
	return all
}
