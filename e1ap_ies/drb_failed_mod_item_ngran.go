package e1ap_ies

// DRBFailedModItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:505
type DRBFailedModItemNGRAN struct {
	DRBID int64 `asn1:"lb:1,ub:32,mandatory,ext"`
	Cause Cause `asn1:"mandatory,ext"`
}
