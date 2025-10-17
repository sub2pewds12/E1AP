package e1ap_ies

// DRBToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBToModifyItemEUTRAN struct {
	DRBID                     DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration         *PDCPConfiguration          `aper:"optional,ext"`
	EUTRANQOS                 *EUTRANQOS                  `aper:"optional,ext"`
	S1ULUPTNLInformation      *UPTNLInformation           `aper:"optional,ext"`
	DataForwardingInformation *DataForwardingInformation  `aper:"optional,ext"`
	PDCPSNStatusRequest       *PDCPSNStatusRequest        `aper:"optional,ext"`
	PDCPSNStatusInformation   *PDCPSNStatusInformation    `aper:"optional,ext"`
	DLUPParameters            []UPParametersItem          `aper:"optional,ext"`
	CellGroupToAdd            []CellGroupInformationItem  `aper:"optional,ext"`
	CellGroupToModify         []CellGroupInformationItem  `aper:"optional,ext"`
	CellGroupToRemove         []CellGroupInformationItem  `aper:"optional,ext"`
	DRBInactivityTimer        *InactivityTimer            `aper:"lb:1,ub:7200,optional,reject,ext"`
	IEExtensions              *ProtocolExtensionContainer `aper:"optional,ext"`
}
