package crypt

var encryptFactory = make(map[string]Encrypt)

func Register(name string, exec Encrypt) {
	encryptFactory[name] = exec
}

func Load(name string) (Encrypt, error) {
	exec, ok := encryptFactory[name]
	if !ok {
		return nil, ErrInvalidPlugin
	}
	return exec, nil
}

type Encrypt interface {
	Sign(msg []byte, privkey []byte) ([]byte, error)
	Verify(msg []byte, sig []byte, pubkey []byte) (bool, error)
}
