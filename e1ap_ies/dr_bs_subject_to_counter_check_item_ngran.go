package e1ap_ies

// DRBsSubjectToCounterCheckItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:712
type DRBsSubjectToCounterCheckItemNGRAN struct {
	PDUSessionID int64     `asn1:"lb:0,ub:255,mandatory,ext"`
	DRBID        int64     `asn1:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount  PDCPCount `asn1:"mandatory,ext"`
	PDCPDLCount  PDCPCount `asn1:"mandatory,ext"`
}
