package e1ap_ies

// UPParametersItem is a generated SEQUENCE type.
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation            `aper:"mandatory,ext"`
	CellGroupID      CellGroupID                 `aper:"lb:0,ub:3,mandatory,ext"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional,ext"`
	// Possible extensions:
	// - QOSMappingInformation (ID: id-QoS-Mapping-Information)
}
