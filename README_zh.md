<div align="center">
</br>

# KeyLock

| [English](README.md) | 中文 |
| --- | --- |

`KeyLock` 是一个高性能且可扩展的键值锁库，支持针对不同类型键值的并发锁操作。该库提供了一种基于键值的锁机制，允许使用字符串、整数等键值类型进行锁操作，适用于需要高并发和细粒度锁控制的场景。

</div>

## 特性
- 高性能：使用预分配的锁数组减少锁争用。
- 泛型支持：通过引入的泛型特性，实现对不同类型键值的锁支持。
- 扩展性：可以轻松添加对其他类型键值的支持。
- 简单易用：提供统一的接口，便于使用和集成。

## 获取KeyLock

使用 go get 下载和安装 KeyLock：

```bash
go get github.com/lyonnee/keylock
```

## 使用方法

以下是如何使用KeyLock的一个简单例子。

### 字符串锁
```go
func main() {
	// 创建一个字符串键值锁
	textlocker := keylock.NewTextLocker[string](10)

	// 锁定一个字符串键值
	l := textlocker.Lock("exampleKey")
	defer l.Unlock("exampleKey")
}
```

### 数值锁
```go
func main() {
	// 创建一个整数键值锁
	numLocker := keylock.NewNumberLocker[int](512)

	// 锁定一个数值键值
	l := locker.Lock(123)
	defer l.Unlock(123)
}
```

## 有问题?

如果您有任何问题或建议，请随时提交issue。我们非常欢迎大家提供反馈和改进的建议。

## 许可证

此项目遵循MIT许可证。请查阅 [LICENSE文件](./LICENSE) 以获得更多信息。