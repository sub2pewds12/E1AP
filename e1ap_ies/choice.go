package e1ap_ies

// CHOICE From: builtin:-1
// ASN.1 Data Type: CHOICE
const (
	CHOICEPresentNothing uint64 = iota
)

type CHOICE struct {
	Choice uint64
}
