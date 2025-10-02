package e1ap_ies

// GNBCUUPCellGroupRelatedConfiguration From: 9_4_5_Information_Element_Definitions.txt:1086
type GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem

// GNBCUUPCellGroupRelatedConfigurationItem From: 9_4_5_Information_Element_Definitions.txt:1088
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      int64            `asn1:"mandatory"`
	UPTNLInformation UPTNLInformation `asn1:"mandatory"`
	ULConfiguration  *ULConfiguration `asn1:"optional"`
}
