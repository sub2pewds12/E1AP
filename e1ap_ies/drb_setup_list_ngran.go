package e1ap_ies

// DRBSetupListNGRAN From: 9_4_5_Information_Element_Definitions.txt:652
type DRBSetupListNGRAN []DRBSetupItemNGRAN

// DRBSetupItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:654
type DRBSetupItemNGRAN struct {
	DRBID                                int64                      `asn1:"mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem         `asn1:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem              `asn1:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem        `asn1:"optional,ext"`
}
