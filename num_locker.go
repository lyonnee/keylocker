/**
 *@author: Lyon.Nee
 *@date: 2024-05-28
 *@description:
 */

package keylock

import "sync"

type numberLocker[T NumberType] struct {
	s []*sync.Mutex
}

func NewNumberLocker[T NumberType](n int) KeyLocker[T] {
	if n <= 0 {
		n = 512
	}

	locks := make([]*sync.Mutex, n)
	for i := range locks {
		locks[i] = &sync.Mutex{}
	}

	return &numberLocker[T]{
		s: locks,
	}
}

func (l *numberLocker[T]) Lock(key T) *sync.Mutex {
	i := l.getSlot(key)
	l.s[i].Lock()
	return l.s[i]
}

func (l *numberLocker[T]) Unlock(key T) {
	i := l.getSlot(key)
	l.s[i].Unlock()
}

func (l *numberLocker[T]) getSlot(key T) uint32 {
	if key == 0 {
		return 0
	}

	return uint32(key) % uint32(len(l.s))
}
