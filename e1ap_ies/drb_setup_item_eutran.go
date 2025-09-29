package e1ap_ies

// DRBSetupItemEUTRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:623
type DRBSetupItemEUTRAN struct {
	DRBID                             int64                             `asn1:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation                  `asn1:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation        `asn1:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem                `asn1:"lb:1,ub:Item,mandatory,ext"`
	S1DLUPUnchanged                   DRBSetupItemEUTRANS1DLUPUnchanged `asn1:"mandatory,ext"`
}
