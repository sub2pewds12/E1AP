package e1ap_ies

// DRBToSetupItemEUTRAN is a generated SEQUENCE type.
type DRBToSetupItemEUTRAN struct {
	DRBID                            DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `aper:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `aper:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `aper:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `aper:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `aper:"optional,ext"`
	DRBInactivityTimer               *InactivityTimer                  `aper:"lb:1,ub:7200,optional,reject,ext"`
	ExistingAllocatedS1DLUPTNLInfo   *UPTNLInformation                 `aper:"optional,ext"`
	IEExtensions                     *ProtocolExtensionContainer       `aper:"optional,ext"`
}
