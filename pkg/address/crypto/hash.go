package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func Checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}

// Sha2Sum Returns hash: SHA256( SHA256( data ) )
// Where possible, using ShaHash() should be a bit faster
func Sha2Sum(b []byte) []byte {
	tmp := sha256.Sum256(b)
	tmp = sha256.Sum256(tmp[:])
	return tmp[:]
}

// Rimp160 Returns hash: RIMP160( SHA256( data ) )
// Where possible, using RimpHash() should be a bit faster
func Rimp160(b []byte) []byte {
	out := make([]byte, 20)
	rimpHash(b, out[:])
	return out[:]
}

func rimpHash(in []byte, out []byte) {
	sha := sha256.New()
	_, err := sha.Write(in)
	if err != nil {
		return
	}
	rim := ripemd160.New()
	_, err = rim.Write(sha.Sum(nil)[:])
	if err != nil {
		return
	}
	copy(out, rim.Sum(nil))
}
