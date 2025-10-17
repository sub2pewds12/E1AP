package e1ap_ies

// DRBSetupItemNGRAN is a generated SEQUENCE type.
type DRBSetupItemNGRAN struct {
	DRBID                                DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem          `aper:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem               `aper:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem         `aper:"optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer `aper:"optional,ext"`
}
