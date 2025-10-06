package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToRemoveItemEUTRAN is a generated SEQUENCE type.
type DRBRequiredToRemoveItemEUTRAN struct {
	DRBID aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
