## 6.824 Lab 2: Raft

### Introduction
本Lab将实现Raft，一个复制状态机协议。
然后Lab3将基于Raft实现一个KEY/VALUE服务，Lab4将shard服务提升性能。

一个replicated服务通过复制多个副本实现容错，Replication能够保证高可用，即使一些机器crash。
关键的挑战是failure会使得我们的多个副本不一致，也就是数据一致性问题。

Raft会将所有客户端请求组织成序列化所谓的Log，保证所有服务器看到的Log顺序都是一样的。
> Since all the live replicas see the same log contents, they all execute the same requests in the same order, and thus continue to have identical service state. 

也就是说如果一个服务器crash，在它恢复之后，Raft就会帮他看到系统里所有的未处理的Log，
帮助他赶上其他机器。Raft就是这么好的一个管家~

Raft更加强大的地方在于，只要majority的服务器能够正常通信和协作，
系统就能Work。万一系统内已经形成很多分区了，虽然不能make progress，
但是只要整个系统恢复到majority的状态，Raft就能保证整个系统从断掉的地方恢复。

This Lab我们将实现Raft作为一个模块，可以被上层的服务继续调用。
很多Raft实例通过RPC互相通信，来维护Log。
> Your Raft interface will support an indefinite sequence of numbered commands, also called log entries.

要完成的Reading非常多。
- extended Raft paper(**) and the Raft lecture notes
- https://thesquareplanet.com/blog/students-guide-to-raft/
- https://pdos.csail.mit.edu/6.824/labs/raft-locking.txt
- https://pdos.csail.mit.edu/6.824/labs/raft-structure.txt
- For a wider perspective, have a look at Paxos, Chubby, Paxos Made Live, Spanner, Zookeeper, Harp, Viewstamped Replication, and Bolosky et al.


只需要编写raft.go即可，还提供了一个简单的RPC包在labrpc中，可以学习一下。
在src/raft下执行go test，可以看到执行失败的输出：
```
Test (2A): initial election ...
--- FAIL: TestInitialElection2A (5.11s)
    config.go:326: expected one leader, got none
Test (2A): election after network failure ...
--- FAIL: TestReElection2A (5.21s)
    config.go:326: expected one leader, got none
Test (2B): basic agreement ...
--- FAIL: TestBasicAgree2B (10.01s)
    config.go:471: one(100) failed to reach agreement
Test (2B): agreement despite follower disconnection ...
--- FAIL: TestFailAgree2B (10.05s)
    config.go:471: one(101) failed to reach agreement
Test (2B): no agreement if too many followers disconnect ...
```

必须使用labrpc的原因是，labrpc帮我们模拟了延迟、乱序等测试情况，如果我们用网络的话就没办法模拟这些
极端情况了。

最后我们实现的Raft应该暴露如下功能：
```go
// create a new Raft server instance:
rf := Make(peers, me, persister, applyCh)

// start agreement on a new log entry:
rf.Start(command interface{}) (index, term, isleader)

// ask a Raft for its current term, and whether it thinks it is leader
rf.GetState() (term, isLeader)

// each time a new entry is committed to the log, each Raft peer
// should send an ApplyMsg to the service (or tester).
type ApplyMsg
```

### Part 2A
实现Leader选举和心跳（我们知道心跳是AppendEntries附带的功能)，AppendEntries没有log
 entry就是一种形式。
 
2A的目标是能够选举出leader，并且能够替换old leader。

#### hint
- 可以添加需要的state 还需要定义一个log entry的结构体 总之一切遵循论文图2
- Fill in the RequestVoteArgs and RequestVoteReply structs，修改Make方法，开一个go来在后台决定是否要
发起选举。实现RequestVote RPC handler来为他人投票。
- 为了实现心跳，定义AppendEntries结构体，让leader间隔发送。当然我们还需要实现AppendEntries RPC handler来重置
本peer的计时器。
- 保证不同peer不会同一时间发起选举，也就是保证随机时间的重试
- tester需要我们的leader发送心跳不超过一秒10次哦
- tester需要我们的重新选举过程在5s内完成，由于成功选举leader可能需要多轮（split vote等），所以我们心跳
timeout要设小一点，让选举能够尽快发生，尽可能在5s内完成。
- 因为tester限制我们每秒10次心跳，所以我们可以比论文的150-300大一点，但是不要太大，不然不能5s完成。
- Go的rand包会帮忙~
- 我们会有很多goroutine阶段性执行任务，所以我们可以用time.Sleep()来实现。
- 如果失败了，好好读图2。
- 好的debug方式，print大法哈哈哈。go test -run 2A > out 我们打到文件里。
- 注意go rpc 只会发送大写字母开头的field。
- 使用go test -race检测race情况。

### Part 2B
上一部分实现的是选举，这一部分实现log replicate。
主要目标是AppendEntries的处理。

### Part 2C
Writing Raft's persistent state to disk each time it changes。

这里将使用persister.go来帮助我们持久化，不会真的写到磁盘上，只是保存在内存里。

每次状态改变，我们都应该保存下来。Use the Persister's ReadRaftState() and SaveRaftState() methods.






