package e1ap_ies

// DRBsSubjectToCounterCheckItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:698
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID       int64     `asn1:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount PDCPCount `asn1:"mandatory,ext"`
	PDCPDLCount PDCPCount `asn1:"mandatory,ext"`
}
