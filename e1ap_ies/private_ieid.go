package e1ap_ies

// PrivateIEID is a generated CHOICE type.
type PrivateIEID struct {
	Choice uint64
	Local  *PrivateIEIDLocal
	Global *string
}

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)
