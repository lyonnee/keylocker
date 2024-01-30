<div align="center">
</br>

# KeyLock

| English | [中文](README.md) |
| --- | --- |

KeyLock is a concurrency-safe key lock manager library written in Go. It can be used to create and manage `ReadWriteLocks`, and dynamically obtain the lock needed through the key.

</div>

## Get KeyLock

You can download KeyLock to your project with the following command.

```bash
go get github.com/lyonnee/keylock
```

## Usage

Here is a simple example of how to use KeyLock.

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lyonnee/keylock"
)

func main() {
    // Create a KeyLocker with a length of 512.
	keylocker := keylock.NewKeyLock(512)

    key := "myKey"

    // Use KeyLocker's Lock method to get the ReadWriteLock corresponding to the key and lock it
	m := keylocker.Lock(key)

    // Example: perform some operations
    fmt.Println("do something...")
    time.Sleep(time.Second)

    // Use Unlock method to release the ReadWriteLock corresponding to the key
	m.Unlock()
}
```

In the above code, we first create a new KeyLocker instance (storing 512 ReadWriteLocks) via `keylocker.NewKeyLock(512)`, then we take out a ReadWriteLock associated with the specified key and lock it via `keylocker.Lock(key)`. After performing necessary operations, we unlock the ReadWriteLock via the `m.Unlock()` method.

At the same time, KeyLocker uses the global lock `sync.Mutex` to ensure concurrency safety, and it uses hash to manage keys. In this way, the ReadWriteLocks corresponding to different keys will not affect each other, making it suitable for scenarios where concurrent processing of multiple tasks is required.

## Questions?

If you have any questions or suggestions, please feel free to submit an issue. We appreciate any feedback and improvement suggestions.

## License

This project follows the MIT license. Please refer to the [LICENSE file](./LICENSE) for more information.