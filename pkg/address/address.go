package address

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/txchat/dtalk/pkg/address/base58"
	"github.com/txchat/dtalk/pkg/address/crypto"
)

//Address 地址
type Address struct {
	Version  byte
	Hash160  [20]byte // For a stealth address: it's HASH160
	Checksum []byte   // Unused for a stealth address
	Pubkey   []byte   // Unused for a stealth address
	Enc58str string
}

//SetBytes 设置地址的bytes
func (a *Address) SetBytes(b []byte) {
	copy(a.Hash160[:], b)
}

func (a *Address) String() string {
	if a.Enc58str == "" {
		var ad [25]byte
		ad[0] = a.Version
		copy(ad[1:21], a.Hash160[:])
		if a.Checksum == nil {
			sh := crypto.Sha2Sum(ad[0:21])
			a.Checksum = make([]byte, 4)
			copy(a.Checksum, sh[:4])
		}
		copy(ad[21:25], a.Checksum[:])
		a.Enc58str = base58.Encode(ad[:])
	}
	return a.Enc58str
}

//NormalVer 普通地址的版本号
const NormalVer byte = 0

func PublicKeyToAddress(version byte, in []byte) string {
	a := new(Address)
	a.Pubkey = make([]byte, len(in))
	copy(a.Pubkey[:], in[:])
	a.Version = version
	a.SetBytes(crypto.Rimp160(in))
	return a.String()
}

//CheckAddress 检查地址
func CheckAddress(ver byte, addr string) (e error) {
	dec := base58.Decode(addr)
	if dec == nil {
		e = errors.New("Cannot decode b58 string '" + addr + "'")
		return
	}
	if len(dec) < 25 {
		e = errors.New("Address too short " + hex.EncodeToString(dec))
		return
	}
	//version 的错误优先
	if dec[0] != ver {
		e = ErrCheckVersion
		return
	}
	//需要兼容以前的错误(以前的错误，是一种特殊的情况)
	if len(dec) == 25 {
		sh := crypto.Sha2Sum(dec[0:21])
		if !bytes.Equal(sh[:4], dec[21:25]) {
			e = ErrCheckChecksum
			return
		}
	}
	var cksum [4]byte
	copy(cksum[:], dec[len(dec)-4:])
	//新的错误: 这个错误用一种新的错误标记
	if crypto.Checksum(dec[:len(dec)-4]) != cksum {
		e = ErrAddressChecksum
	}
	return e
}
