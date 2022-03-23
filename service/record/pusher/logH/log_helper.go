package logH

//用于建立通信层协议seq和业务协议logId的映射关系
type LogHelper struct {
	engine Engine
}

func NewLogHelper(e Engine) *LogHelper {
	return &LogHelper{engine: e}
}

func (lh *LogHelper) sessionToSeq(session string) int32 {
	//TODO
	return 0
}

func (lh *LogHelper) GetLogsIndex(key string, seq int32) (*ConnSeqItem, error) {
	item, err := lh.engine.GetConnSeqIndex(key, seq)
	if err != nil {
		return nil, err
	}
	if item == nil {
		//查询group seq
		session, err := lh.engine.GetGroupSession(key, seq)
		if err != nil {
			return nil, err
		}
		item, err = lh.engine.GetConnSeqIndex(key, lh.sessionToSeq(session))
		if err != nil {
			return nil, err
		}
	}
	return item, nil
}

func (lh *LogHelper) Save(key string, seq int32, item *ConnSeqItem) error {
	return lh.engine.AddConnSeqIndex(key, seq, item)
}

func (lh *LogHelper) SaveProxy(key string, seq int32, item *ConnSeqItem) error {
	//return lh.engine.AddConnSeqIndex(key, seq, item)
	//TODO 实现代理记录索引
	return nil
}
