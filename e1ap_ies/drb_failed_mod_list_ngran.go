package e1ap_ies

// DRBFailedModListNGRAN From: 9_4_5_Information_Element_Definitions.txt:503
type DRBFailedModListNGRAN []DRBFailedModItemNGRAN

// DRBFailedModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:505
type DRBFailedModItemNGRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
