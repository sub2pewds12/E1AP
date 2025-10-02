package e1ap_ies

// DRBSetupModListNGRAN From: 9_4_5_Information_Element_Definitions.txt:668
type DRBSetupModListNGRAN []DRBSetupModItemNGRAN

// DRBSetupModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:670
type DRBSetupModItemNGRAN struct {
	DRBID                                int64                      `asn1:"mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem         `asn1:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem              `asn1:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem        `asn1:"optional,ext"`
}
