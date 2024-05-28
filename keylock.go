/**
 *@author: Lyon.Nee
 *@date: 2024-05-28
 *@description:
 */

package keylock

import "sync"

type NumberType interface {
	~int | ~uint
}

type TextType interface {
	~[]byte | ~string
}

type KeyType interface {
	TextType | NumberType
}

type KeyLocker[T KeyType] interface {
	Lock(key T) *sync.Mutex
	Unlock(key T)
}
