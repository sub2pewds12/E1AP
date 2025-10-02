package e1ap_ies

// DRBToSetupModListNGRAN From: 9_4_5_Information_Element_Definitions.txt:898
type DRBToSetupModListNGRAN []DRBToSetupModItemNGRAN

// DRBToSetupModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:900
type DRBToSetupModItemNGRAN struct {
	DRBID                               int64                             `asn1:"mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `asn1:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `asn1:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `asn1:"mandatory,ext"`
	FlowMappingInformation              []QOSFlowQOSParameterItem         `asn1:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	DRBInactivityTimer                  *int64                            `asn1:"optional,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `asn1:"optional,ext"`
}
