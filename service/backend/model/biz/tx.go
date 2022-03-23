package biz

type TxResult struct {
	Success bool
	To      string
	Amount  int64
	Symbol  string
}
