/**
 *@author: Lyon.Nee
 *@date: 2024-05-28
 *@description:
 */

package keylock

import (
	"hash/fnv"
	"sync"
)

type textLocker[T TextType] struct {
	s []*sync.Mutex
}

func NewTextLocker[T TextType](n int) KeyLocker[T] {
	if n <= 0 {
		n = 512
	}

	locks := make([]*sync.Mutex, n)
	for i := range locks {
		locks[i] = &sync.Mutex{}
	}

	return &textLocker[T]{
		s: locks,
	}
}

func (l *textLocker[T]) Lock(key T) *sync.Mutex {
	i := l.getSlot(key)
	l.s[i].Lock()
	return l.s[i]
}

func (l *textLocker[T]) Unlock(key T) {
	i := l.getSlot(key)
	l.s[i].Unlock()
}

func (l *textLocker[T]) hash(id T) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32()
}

func (l *textLocker[T]) getSlot(key T) uint32 {
	return l.hash(key) % uint32(len(l.s))
}
