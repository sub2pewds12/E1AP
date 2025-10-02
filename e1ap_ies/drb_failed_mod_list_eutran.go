package e1ap_ies

// DRBFailedModListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:477
type DRBFailedModListEUTRAN []DRBFailedModItemEUTRAN

// DRBFailedModItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:479
type DRBFailedModItemEUTRAN struct {
	DRBID int64 `asn1:"mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
