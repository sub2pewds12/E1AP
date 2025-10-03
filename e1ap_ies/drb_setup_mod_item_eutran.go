package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBSetupModItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:639
// ASN.1 Data Type: SEQUENCE
type DRBSetupModItemEUTRAN struct {
	DRBID                             aper.Integer               `aper:"mandatory,ext"`
	S1DLUPTNLInformation              UPTNLInformation           `aper:"mandatory,ext"`
	DataForwardingInformationResponse *DataForwardingInformation `aper:"optional,ext"`
	ULUPTransportParameters           []UPParametersItem         `aper:"mandatory,ext"`
}
