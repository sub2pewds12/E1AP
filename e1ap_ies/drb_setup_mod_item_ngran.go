package e1ap_ies

// DRBSetupModItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:670
type DRBSetupModItemNGRAN struct {
	DRBID                                int64                      `asn1:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation `asn1:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem         `asn1:"lb:1,ub:Item,mandatory,ext"`
	FlowSetupList                        []QOSFlowItem              `asn1:"lb:1,ub:Item,mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem        `asn1:"lb:1,ub:Item,optional,ext"`
}
