package e1ap_ies

// UPParameters From: 9_4_5_Information_Element_Definitions.txt:2396
type UPParameters []UPParametersItem

// UPParametersItem From: 9_4_5_Information_Element_Definitions.txt:2398
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation `asn1:"mandatory,ext"`
	CellGroupID      int64            `asn1:"mandatory,ext"`
}
