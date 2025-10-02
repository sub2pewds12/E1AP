package e1ap_ies

// DRBToRemoveListNGRAN From: 9_4_5_Information_Element_Definitions.txt:813
type DRBToRemoveListNGRAN []DRBToRemoveItemNGRAN

// DRBToRemoveItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:815
type DRBToRemoveItemNGRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
}
