package e1ap_ies

// DRBModifiedItemNGRAN is a generated SEQUENCE type.
type DRBModifiedItemNGRAN struct {
	DRBID                   DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	ULUPTransportParameters []UPParametersItem          `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation    `aper:"optional,ext"`
	FlowSetupList           []QOSFlowItem               `aper:"optional,ext"`
	FlowFailedList          []QOSFlowFailedItem         `aper:"optional,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}
