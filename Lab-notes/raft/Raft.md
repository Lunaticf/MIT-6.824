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



