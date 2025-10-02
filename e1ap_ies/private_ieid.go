package e1ap_ies

// PrivateIEID From: 9_4_6_Common_Definitions.txt:38
const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Choice uint64
	Local  *int64
	Global *string
}
