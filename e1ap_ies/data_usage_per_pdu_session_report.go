package e1ap_ies

// DataUsagePerPDUSessionReport From: 9_4_5_Information_Element_Definitions.txt:346
// ASN.1 Data Type: SEQUENCE
type DataUsagePerPDUSessionReport struct {
	SecondaryRATType          DataUsagePerPDUSessionReportSecondaryRATType `aper:"mandatory,ext"`
	PDUSessionTimedReportList SEQUENCE                                     `aper:"mandatory,ext"`
}
