package record

func (l *Logic) Push(key, from string, body []byte) (int64, uint64, error) {
	return l.svcCtx.AnswerClient.PushCommonMsg(l.ctx, key, from, body)
}
