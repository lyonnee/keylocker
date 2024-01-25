<div align="center">
</br>

# KeyLocker

| [English](README.md) | 中文 |
| --- | --- |

KeyLocker是一个使用Go语言编写的并发安全的键锁管理器库。它可以用来创建和管理互斥锁，并通过键来动态获取当前需要的锁。

</div>

## 获取KeyLocker

您可以通过以下命令将KeyLocker下载到您的项目中。

```bash
go get github.com/lyonnee/keylock
```

## 使用方法

以下是如何使用KeyLocker的一个简单例子。

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
	keyLocker := keylock.NewKeyLock(512)

    key := "myKey"

    // 使用KeyLocker的Lock方法获取对应键的互斥锁并锁定
	m := keyLocker.Lock(key)

    // 举例说明：做一些操作
    fmt.Println("do something...")
    time.Sleep(time.Second)

    // 使用Unlock方法释放对应键的互斥锁
	m.Unlock()
}
```

在上述代码中，我们首先通过`keylock.NewKeyLock(512)`创建一个新的KeyLocker实例(存储512个互斥锁），之后通过`keyLocker.Lock(key)`取出了一个与指定键关联的互斥锁并锁定了它。进行必要的操作后，我们通过`m.Unlock()`方法来解锁该互斥锁。

请注意，您不应在`Unlock`之后再次使用同一个互斥锁，因为在`Unlock`中，该互斥锁可能已经被其他goroutine重用。

同时，KeyLocker由全局锁`sync.Mutex`进行管理，采用了hash的方式对键进行管理，所以不同键对应的互斥锁不会产生影响，非常适合在需要并发处理多任务时使用。

## 有问题?

如果您有任何问题或建议，请随时提交issue。我们非常欢迎大家提供反馈和改进的建议。

## 许可证

此项目遵循MIT许可证。请查阅 [LICENSE文件](./LICENSE) 以获得更多信息。