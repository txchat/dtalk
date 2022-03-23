package tx

const (
	minFeeRate = 100000
	noFeeRate  = -1
)

// Config 上链服务相关配置
type Config struct {
	Chain   Chain
	Grpc    Grpc
	Encrypt Encrypt
}

type Chain struct {
	FeePrikey string // 代扣私钥字符串
	FeeAddr   string // 代扣地址
	Title     string // 虚拟平行链
	BaseExec  string `json:",default=none"`  // 代扣执行器
	CoinExec  string `json:",default=coins"` // 平行链币消耗执行器
	TokenExec string `json:",default=token"` // 平行链币消耗执行器
	Coin      string // 需要额外消耗的coin

	ReceiveCoinAddr        string // 收费地址
	ReceiveCoinAmount      int64  `json:",default=100000000"` // 收费数量
	BwalletReceiveCoinAddr string // 接收币钱包支付的地址
}

type Grpc struct {
	BlockChainAddr string
}

type Encrypt struct {
	Seed     string
	SignType int32 `json:",default=1"` // 交易签名方式 	Invalid=0,SECP256K1=1,ED25519=2,SM2=3
}

type TxInfo struct {
	Exec    string
	Payload []byte
	FeeRate int64
	Prikey  string
}

// 区块链上链的option字段
type BlockProofOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BlockProof struct {
	Data    string `json:"data"`
	Version string `json:"version"`
	Option  string `json:"option"`
	Note    string `json:"note"`
	Ext     string `json:"ext"`
}

type BlockUserConfig struct {
	Op               string `json:"op"`
	Organization     string `json:"organization"`
	Role             string `json:"role"`
	Address          string `json:"address"`
	MemberNote       string `json:"member_note"`
	OrganizationNote string `json:"organization_note"`
}

// BlockDeleteProof 上链数据格式
type BlockDeleteProof struct {
	Id    string `json:"id"` // 存证哈希
	Note  string `json:"note"`
	Force bool   `json:"force"`
}

// 模板上链格式
type BlockTemplate struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type QueryProof struct {
	Id     int64
	Hash   string
	IsBase bool
}

type ReqChainProofs struct {
	Hashs []string `json:"hash"`
}

type RespChainProofs struct {
	BaseHash          string `json:"basehash"`
	PreHash           string `json:"prehash"`
	ProofBlockHash    string `json:"proof_block_hash"`
	ProofBlockTime    int64  `json:"proof_block_time"`
	ProofData         string `json:"proof_data"`
	ProofDeleted      string `json:"proof_deleted"`
	ProofDeletedNote  string `json:"proof_deleted_note"`
	ProofHeight       int64  `json:"proof_height"`
	ProofHeightIndex  int64  `json:"proof_height_index"`
	ProofId           string `json:"proof_id"`
	ProofNote         string `json:"proof_note"`
	ProofOrganization string `json:"proof_organization"`
	ProofSender       string `json:"proof_sender"`
	ProofTxHash       string `json:"proof_tx_hash"`
}

type ReqChainProofMember struct {
	Address []string
}

type RespChainProofMember struct {
	Address      string `json:"address"`
	Role         string `json:"role"`
	Organization string `json:"organization"`
	Note         string `json:"note"`
	Height       int64  `json:"height"`
	Ts           int64  `json:"ts"`
	BlockHash    string `json:"block_hash"`
	Index        int64  `json:"index"`
	Send         string `json:"send"`
	TxHash       string `json:"tx_hash"`
	HeightIndex  int64  `json:"height_index"`
}

type GenAddressAndPrikeyRes struct {
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
	Prikey  string `json:"prikey"`
}
