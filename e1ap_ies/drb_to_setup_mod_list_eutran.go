package e1ap_ies

// DRBToSetupModListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:858
type DRBToSetupModListEUTRAN []DRBToSetupModItemEUTRAN

// DRBToSetupModItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:860
type DRBToSetupModItemEUTRAN struct {
	DRBID                            int64                             `asn1:"mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `asn1:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `asn1:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `asn1:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `asn1:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `asn1:"optional,ext"`
	DRBInactivityTimer               *int64                            `asn1:"optional,ext"`
}
