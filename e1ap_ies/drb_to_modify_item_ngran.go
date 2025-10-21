package e1ap_ies

// DRBToModifyItemNGRAN is a generated SEQUENCE type.
type DRBToModifyItemNGRAN struct {
	DRBID                        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration            *SDAPConfiguration          `aper:"optional,ext"`
	PDCPConfiguration            *PDCPConfiguration          `aper:"optional,ext"`
	DRBDataForwardingInformation *DataForwardingInformation  `aper:"optional,ext"`
	PDCPSNStatusRequest          *PDCPSNStatusRequest        `aper:"optional,ext"`
	PdcpSNStatusInformation      *PDCPSNStatusInformation    `aper:"optional,ext"`
	DLUPParameters               []UPParametersItem          `aper:"optional,ext"`
	CellGroupToAdd               []CellGroupInformationItem  `aper:"optional,ext"`
	CellGroupToModify            []CellGroupInformationItem  `aper:"optional,ext"`
	CellGroupToRemove            []CellGroupInformationItem  `aper:"optional,ext"`
	FlowMappingInformation       []QOSFlowQOSParameterItem   `aper:"optional,ext"`
	DRBInactivityTimer           *InactivityTimer            `aper:"lb:1,ub:7200,optional,reject,ext"`
	IEExtensions                 *ProtocolExtensionContainer `aper:"optional,ext"`
	// Possible extensions:
	// - QOSFlowList (ID: id-OldQoSFlowMap-ULendmarkerexpected)
	// - QoSFlowLevelQoSParameters (ID: id-DRB-QoS)
	// - EarlyForwardingCOUNTReq (ID: id-EarlyForwardingCOUNTReq)
	// - EarlyForwardingCOUNTInfo (ID: id-EarlyForwardingCOUNTInfo)
	// - DAPSRequestInfo (ID: id-DAPSRequestInfo)
	// - EarlyDataForwardingIndicator (ID: id-EarlyDataForwardingIndicator)
}
