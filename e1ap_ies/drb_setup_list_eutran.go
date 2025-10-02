package e1ap_ies

// DRBSetupListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:621
type DRBSetupListEUTRAN []DRBSetupItemEUTRAN

// DRBSetupItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:623
type DRBSetupItemEUTRAN struct {
	DRBID                             int64                              `asn1:"mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation                   `asn1:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation         `asn1:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem                 `asn1:"mandatory,ext"`
	S1DLUPUnchanged                   *DRBSetupItemEUTRANS1DLUPUnchanged `asn1:"optional,ext"`
}
