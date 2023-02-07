package authmock

var kvStore = map[string]string{
	"MOCK":  "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
	"MOCK2": "1AsPsahP7FvpR7F2de1LhSB4SU5ShqZ7eu",
}

type KVMock struct {
}

func NewKVMock() *KVMock {
	return &KVMock{}
}

func (m *KVMock) Signature(sig string) string {
	return kvStore[sig]
}
