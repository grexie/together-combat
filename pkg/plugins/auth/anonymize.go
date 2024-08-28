package auth

type AppAnonymization interface {
	Key() []byte
	SetKey(key []byte)
	WithKey(key []byte) AppAnonymization
}