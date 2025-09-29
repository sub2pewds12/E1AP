package e1ap_ies

// DRBActivityItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:427
type DRBActivityItem struct {
	DRBID       int64       `asn1:"lb:1,ub:32,mandatory,ext"`
	DRBActivity DRBActivity `asn1:"mandatory,ext"`
}
