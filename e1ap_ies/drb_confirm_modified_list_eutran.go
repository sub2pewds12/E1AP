package e1ap_ies

// DRBConfirmModifiedListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:438
type DRBConfirmModifiedListEUTRAN []DRBConfirmModifiedItemEUTRAN

// DRBConfirmModifiedItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:440
type DRBConfirmModifiedItemEUTRAN struct {
	DRBID                int64                      `asn1:"mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `asn1:"optional,ext"`
}
