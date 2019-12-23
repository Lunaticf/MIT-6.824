## Bitcoin FAQ
https://pdos.csail.mit.edu/6.824/papers/bitcoin-faq.txt

Q: It takes an average of 10 minutes for a Bitcoin block to be
validated. Does this mean that the parties involved aren't sure if the
transaction really happened until 10 minutes later?

小额交易大可不必，只需要查询下全节点该笔交易有没有被花过即可。
大额交易，应当等到足够多的区块确认。

Q: What can be done to speed up transactions on the blockchain?

forks are bad，因为产生了不一致。 如果flooding time可以减小，10min就可以减小，就提高了TPS。

Q: The entire blockchain needs to be downloaded before a node can
participate in the network. Won't that take an impractically long time
as the blockchain grows?

现在这也是个问题，区块链体积爆炸。但是全节点毕竟少数。


Q: From some news stories, I have heard that a large number of bitcoin
miners are controlled by a small number of companies.

https://blockchain.info/pools

三个矿池占据了超过51%的算力。

Q: Are there any ways for Bitcoin mining to do useful work, beyond simply
brute-force calculating SHA-256 hashes?

我们可以不做那些无意义的Hash运算吗？做一些有意义的事。
https://www.cs.umd.edu/~elaine/docs/permacoin.pdf

用来算素数。

Q: There is hardware specifically designed to mine Bitcoin. How does
this type of hardware differ from the type of hardware in a laptop?

现在矿机都是专有硬件，对SHA256运算友好。

Q: The paper estimates that the disk space required to store the block
chain will by 4.2 megabytes per year. That seems very low!

只是区块头啦。

Q: Bitcoin uses the hash of the transaction record to identify the
transaction, so it can be named in future transactions. Is this
guaranteed to lead to unique IDs?

其实每笔交易并不是有唯一的id表示，只不过碰撞的概率非常低啦。


Q: Satoshi's paper mentions that each transaction has its own
transaction fees that are given to whoever mined the block. Why would
a miner not simply try to mine blocks with transactions with the
highest transaction fees?

为什么矿工不只打包高交易费的交易？
实际上就是这样的，矿工当然愿意打包高交易非的交易。
https://en.bitcoin.it/wiki/Transaction_fees

Q: How is the price of Bitcoin determined?

当然是供需关系啦。