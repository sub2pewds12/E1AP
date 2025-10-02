package e1ap_ies

// DataUsagePerPDUSessionReport From: 9_4_5_Information_Element_Definitions.txt:346
type DataUsagePerPDUSessionReport struct {
	SecondaryRATType          DataUsagePerPDUSessionReportSecondaryRATType `asn1:"mandatory,ext"`
	PDUSessionTimedReportList SEQUENCE                                     `asn1:"mandatory,ext"`
}
