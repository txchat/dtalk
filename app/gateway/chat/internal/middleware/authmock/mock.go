package authmock

type Mock interface {
	Signature(sig string) string
}
