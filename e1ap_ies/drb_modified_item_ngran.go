package e1ap_ies

// DRBModifiedItemNGRAN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:561
type DRBModifiedItemNGRAN struct {
	DRBID                   int64                    `asn1:"lb:1,ub:32,mandatory,ext"`
	ULUPTransportParameters []UPParametersItem       `asn1:"lb:1,ub:Item,optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `asn1:"optional,ext"`
	FlowSetupList           []QOSFlowItem            `asn1:"lb:1,ub:Item,optional,ext"`
	FlowFailedList          []QOSFlowFailedItem      `asn1:"lb:1,ub:Item,optional,ext"`
}
