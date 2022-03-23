package address

import "errors"

// ErrCheckVersion :
var ErrCheckVersion = errors.New("check version error")

//ErrCheckChecksum :
var ErrCheckChecksum = errors.New("Address Checksum error")

//ErrAddressChecksum :
var ErrAddressChecksum = errors.New("address checksum error")
