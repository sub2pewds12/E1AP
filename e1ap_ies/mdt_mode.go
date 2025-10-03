package e1ap_ies

// MDTMode From: 9_4_5_Information_Element_Definitions.txt:1394
// ASN.1 Data Type: CHOICE
const (
	MDTModePresentNothing uint64 = iota
	MDTModePresentImmediateMDT
)

type MDTMode struct {
	Choice       uint64
	ImmediateMDT *ImmediateMDT
}
