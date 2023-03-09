package secp256K1

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var (
	mnemonic   = "游 即 暗 体 柬 京 非 李 限 稻 跳 务 桥 凶 溶"
	privateKey = "fa884fd1b47d9e9e8dc19e47dc1a794a524ce5d4ee1b82ec92b1ffc1f109c2b6"
	publicKey  = "022db0e08669b30c5dab8c564b428db4944912144088943ec9b690a9046bc8f78b"
	address    = "1AqutxNoVTtcWiVYpBtvficAgea1dYTddR"

	message = "hello world"
)

func Test_ComparePublicKey(t *testing.T) {
	public, err := hex.DecodeString(publicKey)
	assert.Nil(t, err)

	ecdsaPubKey, err := crypto.DecompressPubkey(public)
	assert.Nil(t, err)

	ecdsaPubKeyByte := crypto.FromECDSAPub(ecdsaPubKey)
	t.Log("decompressed public key:", hex.EncodeToString(ecdsaPubKeyByte))
	//t.Log("65 byte hybrid format public key:", hex.EncodeToString(btcecPubKey.SerializeHybrid()))
	t.Log("compressed public key:", hex.EncodeToString(crypto.CompressPubkey(ecdsaPubKey)))

	pub, err := crypto.UnmarshalPubkey(ecdsaPubKeyByte)
	assert.Nil(t, err)
	assert.Equal(t, true, pub.Equal(ecdsaPubKey))
}

// msg: 原文；sigHex：签名后的hex string
func verify(msg, sigHex string) (bool, error) {
	var eth ethereum
	public, err := hex.DecodeString(publicKey)
	if err != nil {
		return false, err
	}
	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return false, err
	}
	msg256 := sha256.Sum256([]byte(msg))
	return eth.Verify(msg256[:], sig, public)
}

func Test_ChatSign(t *testing.T) {
	var eth ethereum
	msg256 := sha256.Sum256([]byte(message))

	private, err := hex.DecodeString(privateKey)
	assert.Nil(t, err)

	sig, err := eth.Sign(msg256[:], private)
	assert.Nil(t, err)

	sigHex := hex.EncodeToString(sig)
	t.Log(sigHex)

	ret, err := verify(message, sigHex)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)
}

func Test_VerifyHaltingstate(t *testing.T) {
	sigHex := "0b7feff8af5cd69c2d20a3dbde6014292cd6d95d8c5e28d3111e9f6e7939108b2926d9dad50cef9c347b53b8fff64be0134beefdf592fe08f069a5d87e8a34ea00"
	ret, err := verify(message, sigHex)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)
}

func Test_VerifyBTC(t *testing.T) {
	sigHex := "b358a084270887a76eed1a9c1ef34c3e47fab875ce50e3074509e1f4e2834b8312af78c34b3769d373e204a4f4b015b69c46e5f7721134b9d8950d2be0e8fdf800"
	ret, err := verify(message, sigHex)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)
}

func Test_VerifyETH(t *testing.T) {
	sigHex := "b358a084270887a76eed1a9c1ef34c3e47fab875ce50e3074509e1f4e2834b8312af78c34b3769d373e204a4f4b015b69c46e5f7721134b9d8950d2be0e8fdf800"
	ret, err := verify(message, sigHex)
	assert.Nil(t, err)
	assert.Equal(t, true, ret)
}
