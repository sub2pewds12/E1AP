package e1ap_ies

// NRCGISupportItem From: 9_4_5_Information_Element_Definitions.txt:1504
// ASN.1 Data Type: SEQUENCE
type NRCGISupportItem struct {
	NRCGI NRCGI `aper:"mandatory"`
}
