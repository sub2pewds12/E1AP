package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBSetupItemNGRAN is a generated SEQUENCE type.
type DRBSetupItemNGRAN struct {
	DRBID                                aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	DRBDataForwardingInformationResponse *DataForwardingInformation  `aper:"optional,ext"`
	ULUPTransportParameters              []UPParametersItem          `aper:"mandatory,ext"`
	FlowSetupList                        []QOSFlowItem               `aper:"mandatory,ext"`
	FlowFailedList                       []QOSFlowFailedItem         `aper:"optional,ext"`
	IEExtensions                         *ProtocolExtensionContainer `aper:"optional,ext"`
}
