package e1ap_ies

// PDUSessionResourceDataUsageItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1597
type PDUSessionResourceDataUsageItem struct {
	PDUSessionID         int64                `asn1:"lb:0,ub:255,mandatory,ext"`
	MRDCUsageInformation MRDCUsageInformation `asn1:"mandatory,ext"`
}
