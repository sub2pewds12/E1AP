package e1ap_ies

// DRBToModifyListNGRAN From: 9_4_5_Information_Element_Definitions.txt:761
type DRBToModifyListNGRAN []DRBToModifyItemNGRAN

// DRBToModifyItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:763
type DRBToModifyItemNGRAN struct {
	DRBID                        int64                      `asn1:"mandatory,ext"`
	SDAPConfiguration            *SDAPConfiguration         `asn1:"optional,ext"`
	PDCPConfiguration            *PDCPConfiguration         `asn1:"optional,ext"`
	DRBDataForwardingInformation *DataForwardingInformation `asn1:"optional,ext"`
	PDCPSNStatusRequest          *PDCPSNStatusRequest       `asn1:"optional,ext"`
	PdcpSNStatusInformation      *PDCPSNStatusInformation   `asn1:"optional,ext"`
	DLUPParameters               []UPParametersItem         `asn1:"optional,ext"`
	CellGroupToAdd               []CellGroupInformationItem `asn1:"optional,ext"`
	CellGroupToModify            []CellGroupInformationItem `asn1:"optional,ext"`
	CellGroupToRemove            []CellGroupInformationItem `asn1:"optional,ext"`
	FlowMappingInformation       []QOSFlowQOSParameterItem  `asn1:"optional,ext"`
	DRBInactivityTimer           *int64                     `asn1:"optional,ext"`
}
