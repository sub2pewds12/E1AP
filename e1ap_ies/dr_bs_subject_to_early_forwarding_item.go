package e1ap_ies

// DRBsSubjectToEarlyForwardingItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:727
type DRBsSubjectToEarlyForwardingItem struct {
	DRBID        int64     `asn1:"lb:1,ub:32,mandatory,ext"`
	DLCountValue PDCPCount `asn1:"mandatory,ext"`
}
