package e1ap_ies

// DRBConfirmModifiedListNGRAN From: 9_4_5_Information_Element_Definitions.txt:451
type DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN

// DRBConfirmModifiedItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:453
type DRBConfirmModifiedItemNGRAN struct {
	DRBID                int64                      `asn1:"mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `asn1:"optional,ext"`
}
