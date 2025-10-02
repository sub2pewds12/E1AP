package e1ap_ies

// DRBToSetupListNGRAN From: 9_4_5_Information_Element_Definitions.txt:877
type DRBToSetupListNGRAN []DRBToSetupItemNGRAN

// DRBToSetupItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:879
type DRBToSetupItemNGRAN struct {
	DRBID                               int64                             `asn1:"mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `asn1:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `asn1:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `asn1:"mandatory,ext"`
	QOSFlowInformationToBeSetup         []QOSFlowQOSParameterItem         `asn1:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	DRBInactivityTimer                  *int64                            `asn1:"optional,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `asn1:"optional,ext"`
}
