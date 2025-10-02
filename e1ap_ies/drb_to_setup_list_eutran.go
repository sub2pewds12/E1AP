package e1ap_ies

// DRBToSetupListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:838
type DRBToSetupListEUTRAN []DRBToSetupItemEUTRAN

// DRBToSetupItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:840
type DRBToSetupItemEUTRAN struct {
	DRBID                            int64                             `asn1:"mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `asn1:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `asn1:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `asn1:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `asn1:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `asn1:"optional,ext"`
	DRBInactivityTimer               *int64                            `asn1:"optional,ext"`
	ExistingAllocatedS1DLUPTNLInfo   *UPTNLInformation                 `asn1:"optional,ext"`
}
