<div align="center">
</br>

# KeyLock

| English | [中文](README_zh.md) |
| --- | --- |

`KeyLock` is a high-performance and scalable key-value lock library that supports concurrent lock operations for different types of keys. This library provides a key-based locking mechanism, allowing lock operations using key types such as strings and integers, suitable for scenarios that require high concurrency and fine-grained lock control.

</div>

## Features
- High Performance: Uses a pre-allocated array of locks to reduce lock contention.
- Generic Support: Introduces generic features to implement lock support for different types of keys.
- Scalability: Easy to add support for other types of keys.
- Easy to Use: Provides a unified interface for easy use and integration.

## Getting KeyLock

Download and install KeyLock using the go get command:

```bash
go get github.com/lyonnee/keylock
```

## Usage

Here is a simple example of how to use KeyLock.

### String Lock
```go
func main() {
    // Create a string key-value lock
    textLocker := keylock.NewTextLocker[string](10)

    // Lock a string key
    l := textLocker.Lock("exampleKey")
    defer l.Unlock("exampleKey")
}
```

### Numeric Lock
```go
func main() {
    // Create an integer key-value lock
    numLocker := keylock.NewNumberLocker[int](512)

    // Lock a numeric key
    l := numLocker.Lock(123)
    defer l.Unlock(123)
}
```

## Questions?

If you have any questions or suggestions, please feel free to submit an issue. We appreciate any feedback and improvement suggestions.

## License

This project follows the MIT license. Please refer to the [LICENSE file](./LICENSE) for more information.