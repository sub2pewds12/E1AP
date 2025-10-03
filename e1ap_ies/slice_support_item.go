package e1ap_ies

// SliceSupportItem From: 9_4_5_Information_Element_Definitions.txt:2176
// ASN.1 Data Type: SEQUENCE
type SliceSupportItem struct {
	SNSSAI SNSSAI `aper:"mandatory"`
}
