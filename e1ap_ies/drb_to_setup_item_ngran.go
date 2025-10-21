package e1ap_ies

// DRBToSetupItemNGRAN is a generated SEQUENCE type.
type DRBToSetupItemNGRAN struct {
	DRBID                               DRBID                             `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `aper:"mandatory,ext"`
	QOSFlowInformationToBeSetup         []QOSFlowQOSParameterItem         `aper:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *InactivityTimer                  `aper:"lb:1,ub:7200,optional,reject,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
	IEExtensions                        *ProtocolExtensionContainer       `aper:"optional,ext"`
	// Possible extensions:
	// - QoSFlowLevelQoSParameters (ID: id-DRB-QoS)
	// - DAPSRequestInfo (ID: id-DAPSRequestInfo)
	// - IgnoreMappingRuleIndication (ID: id-ignoreMappingRuleIndication)
	// - QOSFlowsDRBRemapping (ID: id-QoSFlowsDRBRemapping)
}
