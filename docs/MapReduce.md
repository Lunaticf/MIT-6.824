## MapReduce

论文链接：https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf

#### 自然的思考：
要让多个machine执行一个任务：
1. 如何分解任务
2. 分解好之后，如何分配任务
3. 如何汇总结果

#### 细化该模型：

Data-存在各个机器上

协调
- 需要一个master分配、汇总和协调，多个worker
- worker分为mapper和reducer，mapper发送结果到reducer，reducer汇总，再到master
- 一个reducer处理多个mapper的output，一个mapper会发送多个结果到不同的reducer 多对多的relationship
 
#### Example-词频统计：

有一千万个文档，总共包含2000个单词，那么任务的input=10000000，output=2000
定义100个mapper, 10个reducer，
每个mapper处理10万个文档，将涉及到的单词的词频统计传给reducer，
注意，某个单词将会传送到特定的reducer，如不同mapper计算出的dog单词词频都会传送给reducer2，
每个reducer负责一部分单词的汇总。

#### 难点：
- 分配任务
- 数据传送
- 应对failure
- 追踪完成的任务
- ....

#### 设计：
论文写于2004，那时候带宽还是差的比较多的，所以比较关注减少data在网络中的传送，
现在data center速度要快很多，所以如何解决slow network。
Mapper的input都是在本地磁盘上读取的，整个MapReduce中只有shuffle过程产生一次网络传输。
Q: Why not stream the records to the reducer (via TCP) as they are being
       produced by the mappers?
我想一个原因是如果reducer crash了，计算结果就lost了，还不如写完再一起传输。

有的机器可能完成任务更快，不可能让他先完成在那里空闲呀，尽可能的把任务切分的小一点，
在快的机器完成之后再派发更多的任务。

MapReduce的容错模型相对简单，这也是它简洁强有力的原因。
如果一个机器crash，我们只需要rerun那个failure的Map或者Reduce任务即可，
因为每个function都是没有副作用的（Pure），不会依赖于之前的状态，同阶段任务之间也不会有协调。





知乎上有个答案比较好，引用一下，大而化小，异而化同，非常精辟。
MapReduce可以说是一种计算模型，具体应用和实现有很多难点和trade-off。
ps：shuffle不是洗牌的意思么...应该是有序到无序，为啥来形容Mapper和Reducer交互过程这个异而化同的熵减过程呢...

## Lab1
两个模式的MapReduce
sequential和distributed，
sequential方便我们调试，其Mapper和Reduce阶段都是串行执行的。

### 序言 熟悉一下代码
master.go 两个方法
- Distributed()
- Sequential()

写完Sequential的版本后，我们就能熟悉MapReduce基本的流程了，
下面就是要把它变成分布式的版本。
> One of Map/Reduce's biggest selling points is that it can automatically parallelize ordinary sequential code without any extra work by the developer. 

lab在本机用RPC模拟分布式的计算.

Handling worker failures比较容易
如果worker crash，将会因为超时返回false。
我们只需要重新将该任务分配给其他的worker就可以了。
可以不停地执行call，从register里面取worker，
要注意registerChan是unbuffered channel，会堵塞的。







