package e1ap_ies

// DataUsageReportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:373
type DataUsageReportItem struct {
	DRBID              int64                `asn1:"lb:1,ub:32,mandatory,ext"`
	RATType            RATType              `asn1:"mandatory,ext"`
	DRBUsageReportList []DRBUsageReportItem `asn1:"lb:1,ub:Item,mandatory,ext"`
}
