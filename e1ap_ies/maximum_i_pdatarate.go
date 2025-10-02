package e1ap_ies

// MaximumIPdatarate From: 9_4_5_Information_Element_Definitions.txt:1297
type MaximumIPdatarate struct {
	MaxIPrate MaxIPrate `asn1:"mandatory,ext"`
}
