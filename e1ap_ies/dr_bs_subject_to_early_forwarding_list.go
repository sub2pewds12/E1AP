package e1ap_ies

// DRBsSubjectToEarlyForwardingList From: 9_4_5_Information_Element_Definitions.txt:725
type DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem

// DRBsSubjectToEarlyForwardingItem From: 9_4_5_Information_Element_Definitions.txt:727
type DRBsSubjectToEarlyForwardingItem struct {
	DRBID        int64     `asn1:"mandatory,ext"`
	DLCountValue PDCPCount `asn1:"mandatory,ext"`
}
