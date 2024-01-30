package keylock

import (
	"sync"
	"testing"
)

func TestKeyMutex(t *testing.T) {
	keyMutex := New(512)

	var count = 0

	var wg sync.WaitGroup

	var num = 10000

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			keyMutex.Lock("a")
			count += i
			keyMutex.Unlock("a")
		}(i)
	}

	wg.Wait()

	expected := 50005000

	if count != expected {
		t.Fatalf("exptected %d and actual %d", expected, count)
	}
}

func BenchmarkKeyLock(b *testing.B) {
	keyMutex := New(512)

	var wg sync.WaitGroup

	for j := 0; j < 255; j++ {
		k := string((rune)(j))
		var count = 0
		for i := 0; i <= b.N; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				l := keyMutex.Lock(k)
				count += i
				l.Unlock()
			}(i)
		}
	}

	wg.Wait()
}
