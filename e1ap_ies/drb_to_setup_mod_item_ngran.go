package e1ap_ies

// DRBToSetupModItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:900
type DRBToSetupModItemNGRAN struct {
	DRBID                               int64                             `asn1:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `asn1:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `asn1:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `asn1:"lb:1,ub:Item,mandatory,ext"`
	FlowMappingInformation              []QOSFlowQOSParameterItem         `asn1:"lb:1,ub:Item,mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	DRBInactivityTimer                  *int64                            `asn1:"lb:1,ub:7200,optional,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `asn1:"optional,ext"`
}
