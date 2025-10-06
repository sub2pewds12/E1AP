package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBSetupModItemEUTRAN struct {
	DRBID                             aper.Integer               `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation           `aper:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem         `aper:"mandatory,ext"`
}
