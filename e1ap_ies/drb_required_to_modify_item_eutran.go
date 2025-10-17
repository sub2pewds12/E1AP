package e1ap_ies

// DRBRequiredToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemEUTRAN struct {
	DRBID                                DRBID                                      `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation                 *UPTNLInformation                          `aper:"optional,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	Cause                                Cause                                      `aper:"mandatory,ignore,ext"`
	IEExtensions                         *ProtocolExtensionContainer                `aper:"optional,ext"`
}
