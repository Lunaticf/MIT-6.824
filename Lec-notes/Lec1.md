## Introduction

### 什么是分布式系统
- multiple cooperating computers 
- storage for big web sites, MapReduce, peer-to-peer sharing

### Why 分布
- 协调多个物理实体
- 通过isolation实现security
- 通过replica实现容错
- **scale up throughput via parallel CPUs/mem/disk/net**

### But
- 复杂
- partial failure

### Why this course
- 有趣！
- 现在big web sites都heavily use
- 前沿 big unsolved problems
- hands on 自己实现实际的系统

### Main topics
Hide distribution from applications！
- RPC 线程 并发控制
- Performance
- 容错 hide these failures from the application
- 一致性 一致性和性能是敌人，consistency当然需要更多的communication，lead to more cost





