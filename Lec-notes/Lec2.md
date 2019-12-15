## Infrastructure: RPC and threads
Most commonly-asked question: Why Go?

  6.824 used C++ for many years
  C++ worked out well
  but students spent time tracking down pointer and alloc/free bugs
  and there's no very satisfactory C++ RPC package
    
  Go is a bit better than C++ for us
- good support for concurrency (goroutines, channels, &c)
- good support for RPC
- garbage-collected (no use after freeing problems)
- type safe
- threads + GC is particularly attractive!

线程带来的挑战：
- 共享数据 data race 
    - 使用Mutex
    - 避免共享
- 线程同步
    - 如何等待所有线程执行完成 channel waitGroup
- 并发粒度
    - 粗粒度 好用是好用 可是并发度也降低了
    - 细粒度 more races和死锁


What is a crawler?
爬取所有网页 形成一个图

Crawler challenges
- 尽可能增加每秒能爬取的URL
- 每个URL只爬取一次

Serial crawler:
用一个map避免重复就可以 缺点是一次一个

ConcurrentMutex crawler:
每次用一个线程去fetch，多个线程共享fetch map（data race）

示例，包含了三个版本的Crawler，代码写的非常好
[Crawler.go](../Lec-code/Crawler.go)

那我们什么时候用锁，什么时候用Channel呢？
Most problems can be solved in either style

state -- sharing and locks

communication -- channels

waiting for events -- channels

Use Go's race detector:
    https://golang.org/doc/articles/race_detector.html
    go test -race 
    
Remote Procedure Call (RPC)
分布式系统的基石，客户端与服务端的远程通信。

RPC tries to mimic local fn call:
```go
  Client:
    z = fn(x, y)
  Server:
    fn(x, y) {
      compute
      return z
    }

```

然而现实中比这复杂得多。

[![QhlzlV.md.png](https://s2.ax1x.com/2019/12/15/QhlzlV.md.png)](https://imgse.com/i/QhlzlV)
可以看到上面一个通用的模式，
客户端调用一个stub，stub调用底层的RPC库来发送请求，
请求通过网络传递到服务端，服务端RPC库解析后，通过分发器找到对应的Handlers调用计算并处理。

[kv.go](../Lec-code/kv.go)
可以看到Go写一个RPC服务器非常简单。


RPC problem: what to do about failures?
    那么其实还是要回到最重要的failure的问题，比如丢包、网络延迟、服务器错误等

What does a failure look like to the client RPC library?
- 客户端可能根本收不到服务端的响应
- 客户端也不知道服务器是不是受到请求了
    - 或许没收到
    - 或许收到了 发送回复前崩了
    - 发送回复，网络崩了

怎么处理？

最简单的方法：Best effort
超时重传，但是有很多问题
- 比如发了两次Put, 那么如果Put1在Put2之后执行，就不正确了。
- 适合只读的操作 或者重复操作不会引起服务端的错误，比如数据库不会插入已有的记录

更好的方式-Better RPC behavior: "at most once"

理想化：服务端检测到重复请求，直接丢掉重复请求

那么如何检测重复呢？每个请求有独特的XID，客户端需要给每个re-send分配一样的XID。

```
server:
    if seen[xid]:
      // 把已经执行过的操作结果返回
      r = old[xid]
    else
      // 如果没见过 执行操作
      r = handler()
      old[xid] = r
      seen[xid] = true
```

一些细节：

如何保证XID不重复？Big random number，结合独特的客户端IP等方式。

服务器如果把重复请求信息放在内存，崩溃就不行了。所以可能还要flush到磁盘。

Go RPC is a simple form of "at-most-once"

Go RPC不会重复去试，失败就直接返回失败了。

再讨论一下at-most-once和exactly once

at-most-once是指的0次或1次，也就是可能失败。

exactly once是必须成功且仅成功一次，
那么客户端必须重试，然后配合服务端要有容错机制，重复检测。
这在lab3会实现。













