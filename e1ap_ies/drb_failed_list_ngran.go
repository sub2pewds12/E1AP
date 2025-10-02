package e1ap_ies

// DRBFailedListNGRAN From: 9_4_5_Information_Element_Definitions.txt:490
type DRBFailedListNGRAN []DRBFailedItemNGRAN

// DRBFailedItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:492
type DRBFailedItemNGRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
