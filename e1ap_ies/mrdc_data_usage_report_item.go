package e1ap_ies

// MRDCDataUsageReportItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1315
type MRDCDataUsageReportItem struct {
	StartTimeStamp []byte `asn1:"mandatory,ext"`
	EndTimeStamp   []byte `asn1:"mandatory,ext"`
	UsageCountUL   int64  `asn1:"lb:0,ub:18446744073709551615,mandatory,ext"`
	UsageCountDL   int64  `asn1:"lb:0,ub:18446744073709551615,mandatory,ext"`
}
