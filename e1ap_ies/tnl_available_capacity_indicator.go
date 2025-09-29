package e1ap_ies

// TNLAvailableCapacityIndicator represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2233
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   int64 `asn1:"lb:0,ub:16777216,mandatory,ext"`
	DLTNLAvailableCapacity int64 `asn1:"lb:0,ub:100,mandatory,ext"`
	ULTNLOfferedCapacity   int64 `asn1:"lb:0,ub:16777216,mandatory,ext"`
	ULTNLAvailableCapacity int64 `asn1:"lb:0,ub:100,mandatory,ext"`
}
