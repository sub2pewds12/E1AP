package e1ap_ies

// DRBFailedToModifyListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:516
type DRBFailedToModifyListEUTRAN []DRBFailedToModifyItemEUTRAN

// DRBFailedToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:518
type DRBFailedToModifyItemEUTRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
