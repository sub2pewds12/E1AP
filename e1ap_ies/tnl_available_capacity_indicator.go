package e1ap_ies

// TNLAvailableCapacityIndicator From: 9_4_5_Information_Element_Definitions.txt:2233
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   int64 `asn1:"mandatory,ext"`
	DLTNLAvailableCapacity int64 `asn1:"mandatory,ext"`
	ULTNLOfferedCapacity   int64 `asn1:"mandatory,ext"`
	ULTNLAvailableCapacity int64 `asn1:"mandatory,ext"`
}
