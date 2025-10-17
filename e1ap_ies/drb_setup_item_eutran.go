package e1ap_ies

// DRBSetupItemEUTRAN is a generated SEQUENCE type.
type DRBSetupItemEUTRAN struct {
	DRBID                             DRBID                              `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation                   `aper:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation         `aper:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem                 `aper:"mandatory,ext"`
	S1DLUPUnchanged                   *DRBSetupItemEUTRANS1DLUPUnchanged `aper:"optional,ext"`
	IEExtensions                      *ProtocolExtensionContainer        `aper:"optional,ext"`
}
