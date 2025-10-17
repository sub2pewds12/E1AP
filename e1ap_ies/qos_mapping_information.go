package e1ap_ies

// QOSMappingInformation is a generated SEQUENCE type.
type QOSMappingInformation struct {
	Dscp      *QOSMappingInformationDscp      `aper:"lb:6,ub:6,optional,ext"`
	FlowLabel *QOSMappingInformationFlowLabel `aper:"lb:20,ub:20,optional,ext"`
}
