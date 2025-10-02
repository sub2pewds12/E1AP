package e1ap_ies

// DRBFailedToModifyListNGRAN From: 9_4_5_Information_Element_Definitions.txt:529
type DRBFailedToModifyListNGRAN []DRBFailedToModifyItemNGRAN

// DRBFailedToModifyItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:531
type DRBFailedToModifyItemNGRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
