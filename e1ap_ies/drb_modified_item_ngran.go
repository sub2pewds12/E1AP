package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBModifiedItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:561
// ASN.1 Data Type: SEQUENCE
type DRBModifiedItemNGRAN struct {
	DRBID                   aper.Integer             `aper:"mandatory,ext"`
	ULUPTransportParameters []UPParametersItem       `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `aper:"optional,ext"`
	FlowSetupList           []QOSFlowItem            `aper:"optional,ext"`
	FlowFailedList          []QOSFlowFailedItem      `aper:"optional,ext"`
}
