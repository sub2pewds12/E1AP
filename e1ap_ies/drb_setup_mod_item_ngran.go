package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBSetupModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:670
// ASN.1 Data Type: SEQUENCE
type DRBSetupModItemNGRAN struct {
	DRBID                                aper.Integer               `aper:"mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem         `aper:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem              `aper:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem        `aper:"optional,ext"`
}
