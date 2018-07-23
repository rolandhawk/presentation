package identity

// GOGENERATE OMIT
//go:generate stringer -type IDType
// TYPE OMIT

type IDType int

const (
	KTP IDType = iota
	SIM
	Passport
)

// TYPE OMIT
// GOGENERATE OMIT
