package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBModifiedItemNGRAN is a generated SEQUENCE type.
type DRBModifiedItemNGRAN struct {
	DRBID                   aper.Integer             `aper:"lb:1,ub:32,mandatory,ext"`
	ULUPTransportParameters []UPParametersItem       `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation `aper:"optional,ext"`
	FlowSetupList           []QOSFlowItem            `aper:"optional,ext"`
	FlowFailedList          []QOSFlowFailedItem      `aper:"optional,ext"`
}
