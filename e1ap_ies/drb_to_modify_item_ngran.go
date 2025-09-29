package e1ap_ies

// DRBToModifyItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:763
type DRBToModifyItemNGRAN struct {
	DRBID                        int64                      `asn1:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration            *SDAPConfiguration         `asn1:"optional,ext"`
	PDCPConfiguration            *PDCPConfiguration         `asn1:"optional,ext"`
	DRBDataForwardingInformation *DataForwardingInformation `asn1:"optional,ext"`
	PDCPSNStatusRequest          *PDCPSNStatusRequest       `asn1:"optional,ext"`
	PdcpSNStatusInformation      *PDCPSNStatusInformation   `asn1:"optional,ext"`
	DLUPParameters               []UPParametersItem         `asn1:"lb:1,ub:Item,optional,ext"`
	CellGroupToAdd               []CellGroupInformationItem `asn1:"lb:1,ub:Item,optional,ext"`
	CellGroupToModify            []CellGroupInformationItem `asn1:"lb:1,ub:Item,optional,ext"`
	CellGroupToRemove            []CellGroupInformationItem `asn1:"lb:1,ub:Item,optional,ext"`
	FlowMappingInformation       []QOSFlowQOSParameterItem  `asn1:"lb:1,ub:Item,optional,ext"`
	DRBInactivityTimer           *int64                     `asn1:"lb:1,ub:7200,optional,ext"`
}
