package logH

type Engine interface {
	GetConnSeqIndex(cid string, seq int32) (*ConnSeqItem, error)
	AddConnSeqIndex(cid string, seq int32, item *ConnSeqItem) error
	GetGroupSession(cid string, seq int32) (session string, err error)
}

type ConnSeqItem struct {
	Type   string  `json:"type"`
	Sender string  `json:"sender"`
	Client string  `json:"client"`
	Logs   []int64 `json:"logs"`
}
