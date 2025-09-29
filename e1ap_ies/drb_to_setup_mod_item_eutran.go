package e1ap_ies

// DRBToSetupModItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:860
type DRBToSetupModItemEUTRAN struct {
	DRBID                            int64                             `asn1:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `asn1:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `asn1:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `asn1:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `asn1:"lb:1,ub:Item,mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `asn1:"lb:1,ub:Item,optional,ext"`
	DRBInactivityTimer               *int64                            `asn1:"lb:1,ub:7200,optional,ext"`
}
