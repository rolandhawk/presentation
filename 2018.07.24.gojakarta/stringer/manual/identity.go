package identity

// EXAMPLE OMIT

func (i IDType) String() string {
	switch i {
	case KTP:
		return "KTP"
	case SIM:
		return "SIM"
	case Passport:
		return "Passport"
	}
}

// EXAMPLE OMIT
