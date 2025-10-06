package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToRemoveItemNGRAN is a generated SEQUENCE type.
type DRBToRemoveItemNGRAN struct {
	DRBID aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
}
