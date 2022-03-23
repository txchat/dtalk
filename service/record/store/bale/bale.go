package bale

import xproto "github.com/txchat/imparse/proto"

type Item interface {
	Data() []byte
}

type Bale struct {
	ss []xproto.Proto
}

func Load(item Item) {

}
