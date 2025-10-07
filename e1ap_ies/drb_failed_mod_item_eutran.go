package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBFailedModItemEUTRAN is a generated SEQUENCE type.
type DRBFailedModItemEUTRAN struct {
	DRBID        aper.Integer                `aper:"lb:1,ub:32,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
