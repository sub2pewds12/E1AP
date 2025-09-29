package e1ap_ies

// GNBCUUPCellGroupRelatedConfigurationItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1088
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      int64            `asn1:"lb:0,ub:3,mandatory"`
	UPTNLInformation UPTNLInformation `asn1:"mandatory"`
	ULConfiguration  *ULConfiguration `asn1:"optional"`
}
