package recordhelper

// RecordHelper 用于建立通信层协议seq和业务协议msg_id的映射关系
type RecordHelper struct {
	engine Engine
}

func NewRecordHelper(e Engine) *RecordHelper {
	return &RecordHelper{engine: e}
}

func (lh *RecordHelper) sessionToSeq(session string) int32 {
	//TODO session seq mapping
	return 0
}

func (lh *RecordHelper) GetLogsIndex(key string, seq int32) (*ConnSeqItem, error) {
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

func (lh *RecordHelper) Save(key string, seq int32, item *ConnSeqItem) error {
	return lh.engine.AddConnSeqIndex(key, seq, item)
}

func (lh *RecordHelper) SaveProxy(key string, seq int32, item *ConnSeqItem) error {
	//return lh.engine.AddConnSeqIndex(key, seq, item)
	//TODO 实现代理记录索引
	return nil
}
