package e1ap_ies

// DRBRemovedItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:577
type DRBRemovedItem struct {
	DRBID                     int64                              `asn1:"lb:1,ub:32,mandatory,ext"`
	DRBReleasedInSession      DRBRemovedItemDRBReleasedInSession `asn1:"mandatory,ext"`
	DRBAccumulatedSessionTime *[]byte                            `asn1:"optional,ext"`
}
