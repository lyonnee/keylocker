/**
 *@author: Lyon.Nee
 *@date: 20224-01-25
 *@description:
 */
package keylock

import (
	"hash/fnv"
	"sync"
)

type KeyLocker interface {
	Lock(key string) *sync.RWMutex
	Unlock(key string)
	RLock(key string) *sync.RWMutex
	RUnlock(key string)
}

type KeyLock struct {
	// global m
	m sync.Mutex
	s []*sync.RWMutex
}

// the higher the n-value, the smaller the probability of lock collision
func New(n int) KeyLocker {
	if n <= 0 {
		n = 512 // use 4kb memory in 64bit system
	}

	return &KeyLock{
		s: make([]*sync.RWMutex, n),
	}
}

func (l *KeyLock) Lock(key string) *sync.RWMutex {
	// get index of mutex array
	i := l.hash(key) % uint32(len(l.s))

	l.m.Lock()
	// if nil, create a new mutex pointer
	if l.s[i] == nil {
		l.s[i] = &sync.RWMutex{}
	}
	l.m.Unlock()

	// lock for mutex
	l.s[i].Lock()
	return l.s[i]
}

func (l *KeyLock) Unlock(key string) {
	i := l.hash(key) % uint32(len(l.s))
	l.s[i].Unlock()
}

func (l *KeyLock) RLock(key string) *sync.RWMutex {
	// get index of mutex array
	i := l.hash(key) % uint32(len(l.s))

	l.m.Lock()
	// if nil, create a new mutex pointer
	if l.s[i] == nil {
		l.s[i] = &sync.RWMutex{}
	}
	l.m.Unlock()

	// lock for mutex
	l.s[i].RLock()
	return l.s[i]
}

func (l *KeyLock) RUnlock(key string) {
	i := l.hash(key) % uint32(len(l.s))
	l.s[i].RUnlock()
}

func (l *KeyLock) hash(id string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32()
}
