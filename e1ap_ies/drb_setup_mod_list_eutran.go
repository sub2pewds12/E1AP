package e1ap_ies

// DRBSetupModListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:637
type DRBSetupModListEUTRAN []DRBSetupModItemEUTRAN

// DRBSetupModItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:639
type DRBSetupModItemEUTRAN struct {
	DRBID                             int64                      `asn1:"mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation           `asn1:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem         `asn1:"mandatory,ext"`
}
