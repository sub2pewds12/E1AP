package e1ap_ies

// DRBActivityList From: 9_4_5_Information_Element_Definitions.txt:425
type DRBActivityList []DRBActivityItem

// DRBActivityItem From: 9_4_5_Information_Element_Definitions.txt:427
type DRBActivityItem struct {
	DRBID       int64       `asn1:"mandatory,ext"`
	DRBActivity DRBActivity `asn1:"mandatory,ext"`
}
