package e1ap_ies

// HWCapacityIndicator From: 9_4_5_Information_Element_Definitions.txt:1228
type HWCapacityIndicator struct {
	OfferedThroughput   int64 `asn1:"mandatory,ext"`
	AvailableThroughput int64 `asn1:"mandatory,ext"`
}
