package e1ap_ies

// DRBRequiredToRemoveListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:800
type DRBRequiredToRemoveListEUTRAN []DRBRequiredToRemoveItemEUTRAN

// DRBRequiredToRemoveItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:802
type DRBRequiredToRemoveItemEUTRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
