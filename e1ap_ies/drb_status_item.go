package e1ap_ies

// DRBStatusItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:684
type DRBStatusItem struct {
	DRBID       int64      `asn1:"lb:1,ub:32,mandatory,ext"`
	PDCPDLCount *PDCPCount `asn1:"optional,ext"`
	PDCPULCount *PDCPCount `asn1:"optional,ext"`
}
