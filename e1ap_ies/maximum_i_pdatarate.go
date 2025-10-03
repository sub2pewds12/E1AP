package e1ap_ies

// MaximumIPdatarate From: 9_4_5_Information_Element_Definitions.txt:1297
// ASN.1 Data Type: SEQUENCE
type MaximumIPdatarate struct {
	MaxIPrate MaxIPrate `aper:"mandatory,ext"`
}
