package e1ap_ies

// TReorderingTimer From: 9_4_5_Information_Element_Definitions.txt:2304
// ASN.1 Data Type: SEQUENCE
type TReorderingTimer struct {
	TReordering TReordering `aper:"mandatory,ext"`
}
