package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBModifiedItemEUTRAN is a generated SEQUENCE type.
type DRBModifiedItemEUTRAN struct {
	DRBID                   aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation    *UPTNLInformation           `aper:"optional,ext"`
	PDCPSNStatusInformation *PDCPSNStatusInformation    `aper:"optional,ext"`
	ULUPTransportParameters []UPParametersItem          `aper:"optional,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}
