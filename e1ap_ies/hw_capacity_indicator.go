package e1ap_ies

// HWCapacityIndicator represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1228
type HWCapacityIndicator struct {
	OfferedThroughput   int64 `asn1:"lb:1,ub:16777216,mandatory,ext"`
	AvailableThroughput int64 `asn1:"lb:0,ub:100,mandatory,ext"`
}
