<div align="center">
</br>

# KeyLock

| [English](README.md) | 中文 |
| --- | --- |

KeyLock是一个使用Go语言编写的并发安全的键锁管理器库。它可以用来创建和管理`读写锁`，并通过键来动态获取当前需要的锁。

</div>

## 获取KeyLock

您可以通过以下命令将KeyLock下载到您的项目中。

```bash
go get github.com/lyonnee/keylock
```

## 使用方法

以下是如何使用KeyLock的一个简单例子。

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lyonnee/keylock"
)

func main() {
    // 创建一个长度为512的KeyLocker。
	keylocker := keylock.NewKeyLock(512)

    key := "myKey"

    // 使用KeyLocker的Lock方法获取对应键的读写锁并锁定
	m := keylocker.Lock(key)

    // 举例说明：做一些操作
    fmt.Println("do something...")
    time.Sleep(time.Second)

    // 使用Unlock方法释放对应键的读写锁
	m.Unlock()
}
```

在上述代码中，我们首先通过`keylocker.NewKeyLock(512)`创建一个新的KeyLocker实例(存储512个读写锁），之后通过`keylocker.Lock(key)`取出了一个与指定键关联的读写锁并锁定了它。进行必要的操作后，我们通过`m.Unlock()`方法来解锁该读写锁。

同时，KeyLocker使用全局锁`sync.Mutex`来保证并发安全，它采用了哈希的方式来对键进行管理。通过这种方式，不同的键对应的读写锁不会相互影响，这使得它适合于需要并发处理多任务的场景。

## 有问题?

如果您有任何问题或建议，请随时提交issue。我们非常欢迎大家提供反馈和改进的建议。

## 许可证

此项目遵循MIT许可证。请查阅 [LICENSE文件](./LICENSE) 以获得更多信息。