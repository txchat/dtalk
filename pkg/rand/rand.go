package rand

import (
	"crypto/rand"
	mrand "math/rand"
)

var (
	// stdSource standard source string.
	stdSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// stdNumberSource standard number source string.
	stdNumberSource = "0123456789"

	// 16
	std16Source = "0123456789abcdef"
)

// StdSource returns standard source string.
func StdSource() string {
	return stdSource
}

// StdNumberSource returns standard number source string.
func StdNumberSource() string {
	return stdNumberSource
}

func Std16Source() string {
	return std16Source
}

// NewString returns a new random string of the provided length, consisting
// of the standard source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewString(length int) string {
	return NewWithSource(length, StdSource())
}

// NewNumber returns a new random string of the provided length, consisting
// of the standard number source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewNumber(length int) string {
	return NewWithSource(length, StdNumberSource())
}

// NewWithSource returns a new random string of the provided length, consisting
// of the provided source string.
// It panics if source length is wrong (<1 or >256) or rand.Read occurs an error.
func NewWithSource(length int, source string) string {
	if length == 0 {
		return ""
	}
	sl := len(source)
	if sl < 1 || sl > 256 {
		panic("rand: wrong source length")
	}
	rbMax := 255 - (256 % sl)
	b := make([]byte, length)
	r := make([]byte, length+length/2) // storage for random bytes
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic(err)
		}
		for _, rb := range r {
			v := int(rb)
			if v > rbMax { // skip to avoid modulo bias
				continue
			}
			b[i] = source[v%sl]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func NewAESKey256() string {
	return NewWithSource(32, Std16Source())
}

// RandInt returns int pseudo-random number in [min, max).
func RandInt(min, max int) int {
	if min > max {
		panic("math: min cannot be greater than max")
	}
	if min == max {
		return min
	}
	return mrand.Intn(max-min) + min
}
