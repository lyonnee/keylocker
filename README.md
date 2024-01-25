<div align="center">
</br>

# KeyLocker

| English | [中文](README_zh.md) |
| --- | --- |

KeyLocker is a concurrent-safe key lock manager library written in Go. It can be used to create and manage mutexes and dynamically obtain the required locks through keys.

</div>

## Get KeyLocker

You can download KeyLocker to your project with the following command.

```bash
go get github.com/lyonnee/keylock
```

## Usage

Below is a simple example of how to use KeyLocker.

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
	keyLocker := keylock.NewKeyLock(512)

    key := "myKey"

    // Using the Lock method of KeyLocker to get the mutex associated with the key and lock it
	m := keyLocker.Lock(key)

    // Examples: do some operations
    fmt.Println("do something...")
    time.Sleep(time.Second)

    // Unlock the mutex associated with the key using the Unlock method
	m.Unlock()
}
```

In the above code, we first create a new KeyLocker instance (storing 512 mutexes) through `keylock.NewKeyLock(512)`, and then we obtain a mutex associated with a specific key and lock it with `keyLocker.Lock(key)`. After performing necessary operations, we unlock this mutex with `m.Unlock()` method.

Please note that you should not use the same mutex again after `Unlock`, because in `Unlock`, the mutex may have been reused by other goroutines.

At the same time, KeyLocker is managed by the global lock `sync.Mutex`, and adopts the hash method to manage keys, so the mutexes corresponding to different keys will not affect each other, which is very suitable for use when concurrent processing of multiple tasks is required.

## Any questions?

If you have any questions or suggestions, please feel free to submit an issue. We are very welcome everyone to provide feedback and suggestions for improvement.

## License

This project follows the MIT license. Please refer to the [LICENSE file](./LICENSE) for more information."

The above is a simple example of a README file, you can modify and expand it according to your actual situation.