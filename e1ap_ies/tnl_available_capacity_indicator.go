package e1ap_ies

// TNLAvailableCapacityIndicator is a generated SEQUENCE type.
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   TNLAvailableCapacityIndicatorDLTNLOfferedCapacity   `aper:"lb:0,ub:16777216,mandatory,ext"`
	DLTNLAvailableCapacity TNLAvailableCapacityIndicatorDLTNLAvailableCapacity `aper:"lb:0,ub:100,mandatory,ext"`
	ULTNLOfferedCapacity   TNLAvailableCapacityIndicatorULTNLOfferedCapacity   `aper:"lb:0,ub:16777216,mandatory,ext"`
	ULTNLAvailableCapacity TNLAvailableCapacityIndicatorULTNLAvailableCapacity `aper:"lb:0,ub:100,mandatory,ext"`
	IEExtensions           ProtocolExtensionContainer                          `aper:"mandatory,ext"`
}
