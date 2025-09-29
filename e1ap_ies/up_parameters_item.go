package e1ap_ies

// UPParametersItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2398
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation `asn1:"mandatory,ext"`
	CellGroupID      int64            `asn1:"lb:0,ub:3,mandatory,ext"`
}
