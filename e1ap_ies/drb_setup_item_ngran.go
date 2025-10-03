package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBSetupItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:654
// ASN.1 Data Type: SEQUENCE
type DRBSetupItemNGRAN struct {
	DRBID                                aper.Integer               `aper:"mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem         `aper:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem              `aper:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem        `aper:"optional,ext"`
}
