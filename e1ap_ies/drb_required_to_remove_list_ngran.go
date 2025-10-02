package e1ap_ies

// DRBRequiredToRemoveListNGRAN From: 9_4_5_Information_Element_Definitions.txt:825
type DRBRequiredToRemoveListNGRAN []DRBRequiredToRemoveItemNGRAN

// DRBRequiredToRemoveItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:827
type DRBRequiredToRemoveItemNGRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
