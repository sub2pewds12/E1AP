package e1ap_ies

// DRBToRemoveListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:788
type DRBToRemoveListEUTRAN []DRBToRemoveItemEUTRAN

// DRBToRemoveItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:790
type DRBToRemoveItemEUTRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
}
