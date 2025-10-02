package e1ap_ies

// DRBFailedListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:464
type DRBFailedListEUTRAN []DRBFailedItemEUTRAN

// DRBFailedItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:466
type DRBFailedItemEUTRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
