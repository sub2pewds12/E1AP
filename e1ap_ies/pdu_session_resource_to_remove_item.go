package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceToRemoveItem is a generated SEQUENCE type.
type PDUSessionResourceToRemoveItem struct {
	PDUSessionID aper.Integer `aper:"lb:0,ub:255,mandatory,ext"`
}
