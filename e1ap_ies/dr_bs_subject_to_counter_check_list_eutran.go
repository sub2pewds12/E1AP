package e1ap_ies

// DRBsSubjectToCounterCheckListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:696
type DRBsSubjectToCounterCheckListEUTRAN []DRBsSubjectToCounterCheckItemEUTRAN

// DRBsSubjectToCounterCheckItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:698
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID       int64     `asn1:"mandatory,ext"`
	PDCPULCount PDCPCount `asn1:"mandatory,ext"`
	PDCPDLCount PDCPCount `asn1:"mandatory,ext"`
}
