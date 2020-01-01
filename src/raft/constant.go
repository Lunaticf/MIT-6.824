package raft

type peerState int

// peer的三种状态
const (
	Leader peerState = iota
	Follower
	Candidate
)

// 方便调试 打印状态对应的字符串
func (state peerState) String() string {
	switch state {
	case Leader:
		return "Leader"
	case Follower:
		return "Follower"
	default:
		return "Candidate"
	}
}
