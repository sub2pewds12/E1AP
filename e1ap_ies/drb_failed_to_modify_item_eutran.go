package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBFailedToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBFailedToModifyItemEUTRAN struct {
	DRBID aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
