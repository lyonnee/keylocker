/**
 *@author: Lyon.Nee
 *@date: 20224-01-25
 *@description:
 */
package keylocker

import (
	"hash/fnv"
	"sync"
)

type KeyLocker struct {
	// global m
	m sync.Mutex
	s []*sync.Mutex
}

// the higher the n-value, the smaller the probability of lock collision
func NewKeyLock(n int) *KeyLocker {
	if n <= 0 {
		n = 512 // use 4kb memory in 64bit system
	}

	return &KeyLocker{
		s: make([]*sync.Mutex, n),
	}
}

func (locker *KeyLocker) Lock(key string) *sync.Mutex {
	// get index of mutex array
	i := locker.hash(key) % uint32(len(locker.s))

	locker.m.Lock()
	// if nil, create a new mutex pointer
	if locker.s[i] == nil {
		locker.s[i] = &sync.Mutex{}
	}
	locker.m.Unlock()

	// lock for mutex
	locker.s[i].Lock()
	return locker.s[i]
}

func (locker *KeyLocker) Unlock(key string) {
	i := locker.hash(key) % uint32(len(locker.s))
	locker.s[i].Unlock()
}

func (locker *KeyLocker) hash(id string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32()
}
