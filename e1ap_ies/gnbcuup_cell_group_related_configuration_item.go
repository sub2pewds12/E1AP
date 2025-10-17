package e1ap_ies

// GNBCUUPCellGroupRelatedConfigurationItem is a generated SEQUENCE type.
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      CellGroupID                 `aper:"lb:0,ub:3,mandatory"`
	UPTNLInformation UPTNLInformation            `aper:"mandatory"`
	ULConfiguration  *ULConfiguration            `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}
