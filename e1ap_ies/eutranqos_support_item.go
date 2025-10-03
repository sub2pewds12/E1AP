package e1ap_ies

// EUTRANQOSSupportItem From: 9_4_5_Information_Element_Definitions.txt:1044
// ASN.1 Data Type: SEQUENCE
type EUTRANQOSSupportItem struct {
	EUTRANQOS EUTRANQOS `aper:"mandatory"`
}
