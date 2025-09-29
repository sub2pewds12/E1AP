package e1ap_ies

// DRBFailedModItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:479
type DRBFailedModItemEUTRAN struct {
	DRBID int64 `asn1:"lb:1,ub:32,mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
