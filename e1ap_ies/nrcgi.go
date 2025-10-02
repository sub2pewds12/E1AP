package e1ap_ies

// NRCGI From: 9_4_5_Information_Element_Definitions.txt:1492
type NRCGI struct {
	PLMNIdentity   []byte `asn1:"lb:3,ub:3,mandatory"`
	NRCellIdentity []byte `asn1:"lb:36,ub:36,mandatory"`
}
