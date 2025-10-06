package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToRemoveItemEUTRAN is a generated SEQUENCE type.
type DRBToRemoveItemEUTRAN struct {
	DRBID aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
}
