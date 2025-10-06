package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceFailedModItem is a generated SEQUENCE type.
type PDUSessionResourceFailedModItem struct {
	PDUSessionID aper.Integer `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause        `aper:"mandatory,ignore,ext"`
}
