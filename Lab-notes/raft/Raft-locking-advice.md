## Raft Locking Advice

如果你想知道在raft lab中如何使用锁，这里有一些建议。

### Rule 1
无论什么时候有一个被多个goroutine使用的数据，并且至少有一个线程会修改数据。
那么应该使用锁。Go race detector非常善于检测本条规则是否满足。

### Rule 2
无论什么时候代码对共享数据进行一系列的修改，并且其他线程可能会中途查看数据，
那么应该对整个修改序列使用锁。就是临界区的概念。
```go
// An example:
rf.mu.Lock()
rf.currentTerm += 1
rf.state = Candidate
rf.mu.Unlock()
```
### Rule 3
无论什么时候代码对共享数据读写，如果另外一个线程中途对数据进行修改会引发整个过程错误，
应该使用锁。
```go
//An example that could occur in a Raft RPC handler:
rf.mu.Lock()
if args.Term > rf.currentTerm {
rf.currentTerm = args.Term
}
rf.mu.Unlock()
```
Raft需要currentTerm永远递增。如果另一个线程能够在if和修改语句之间修改currentTerm，最后可能会导致currenTerm减少。
因此锁必须whole sequence持有。

### Rule 4
通常我们加锁的时候，要注意持有锁的时候最好不要做什么需要wait的操作，比如读channel，写channel，等待计时器，
或是sleep等。显然这样别的线程就不能make progress了。
另一个原因是避免死锁，想象两个peer互相发送RPC，所有的RPC handlers都需要接收者的锁，那么产生死锁。

所以，等待的code必须首先释放锁，如果不方便，创建一个独立的线程来wait是很有用的操作。


### Rule 5
以下发送请求投票的代码是不正确的，
```go
rf.mu.Lock()
rf.currentTerm += 1
rf.state = Candidate
for <each peer> {
go func() {
  rf.mu.Lock()
  args.Term = rf.currentTerm
  rf.mu.Unlock()
  Call("Raft.RequestVote", &args, ...)
  // handle the reply...
} ()
}
rf.mu.Unlock()
```
上面代码每次新开一个线程来发送请求，这是不正确的。因为args.Term或许
The code sends each RPC in a separate goroutine. It's incorrect
because args.Term may not be the same as the rf.currentTerm at which
the surrounding code decided to become a Candidate. Lots of time may
pass between when the surrounding code creates the goroutine and when
the goroutine reads rf.currentTerm; for example, multiple terms may
come and go, and the peer may no longer be a candidate. One way to fix
this is for the created goroutine to use a copy of rf.currentTerm made
while the outer code holds the lock. Similarly, reply-handling code
after the Call() must re-check all relevant assumptions after
re-acquiring the lock; for example, it should check that
rf.currentTerm hasn't changed since the decision to become a
candidate.