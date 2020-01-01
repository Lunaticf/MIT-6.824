package raft

// 这里定义了Raft.go里用到的各种结构 比如log entry等

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
// 用于返回客户端信息
//
// in Lab 3 you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh; at that point you can add fields to
// ApplyMsg, but set CommandValid to false for these other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int
}

// log entry 定义
type LogEntry struct {
	LogIndex int
	LogTerm  int
	Command  interface{}
}

// RequestVote RPC的参数
// Candidates收集选票 调用其他peer的RequestVote
type RequestVoteArg struct {
	Term         int // Candidate的term
	CandidateId  int // Candidate的peer索引
	LastLogIndex int // Candidate的最后一个log的index
	LastLogTerm  int // Candidate的最后一个log的term
}

// RequestVote RPC的响应结果
type RequestVoteRes struct {
	Term        int  // 用于Candidate更新term
	VoteGranted bool // true说明同意投票
}

// AppendEntries RPC的参数
// 由Leader调用各Peer的AppendEntries RPC
type AppendEntriesArg struct {
	Term,         // Leader的term
	LeaderId,     // Leader的id
	PrevLogIndex, // 前一个log entry的index
	PrevLogTerm,  // 前一个log entry的term
	CommitIndex int    // Leader的commit index
	Entries []LogEntry // log entry集合 如果是心跳 就是空的
}

//  AppendEntries RPC的响应结果
type AppendEntriesRes struct {
	Term    int  // 用于Leader自省
	Success bool // 如果follower PrevLogIndex和PrevLogTerm都匹配
}
