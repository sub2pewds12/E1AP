package e1ap_ies

// DRBModifiedListNGRAN From: 9_4_5_Information_Element_Definitions.txt:559
type DRBModifiedListNGRAN []DRBModifiedItemNGRAN

// DRBModifiedItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:561
type DRBModifiedItemNGRAN struct {
	DRBID                   int64                    `asn1:"mandatory,ext"`
	ULUPTransportParameters []UPParametersItem       `asn1:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `asn1:"optional,ext"`
	FlowSetupList           []QOSFlowItem            `asn1:"optional,ext"`
	FlowFailedList          []QOSFlowFailedItem      `asn1:"optional,ext"`
}
