package e1ap_ies

// DataUsagePerPDUSessionReport represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:346
type DataUsagePerPDUSessionReport struct {
	SecondaryRATType DataUsagePerPDUSessionReportSecondaryRATType `asn1:"mandatory,ext"`
}
