package e1ap_ies

// DRBRequiredToModifyItemNGRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                DRBID                                      `aper:"lb:1,ub:32,mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `aper:"optional,ext"`
	Cause                                Cause                                      `aper:"mandatory,ignore,ext"`
	IEExtensions                         *ProtocolExtensionContainer                `aper:"optional,ext"`
}
