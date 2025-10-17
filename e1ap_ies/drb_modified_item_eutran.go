package e1ap_ies

// DRBModifiedItemEUTRAN is a generated SEQUENCE type.
type DRBModifiedItemEUTRAN struct {
	DRBID                   DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation           `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation    `aper:"optional,ext"`
	ULUPTransportParameters []UPParametersItem          `aper:"optional,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}
