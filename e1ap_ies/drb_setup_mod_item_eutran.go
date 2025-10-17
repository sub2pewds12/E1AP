package e1ap_ies

// DRBSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBSetupModItemEUTRAN struct {
	DRBID                             DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation            `aper:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem          `aper:"mandatory,ext"`
	IEExtensions                      *ProtocolExtensionContainer `aper:"optional,ext"`
}
