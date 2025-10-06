package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceActivityItem is a generated SEQUENCE type.
type PDUSessionResourceActivityItem struct {
	PDUSessionID               aper.Integer               `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionResourceActivity PDUSessionResourceActivity `aper:"mandatory,ext"`
}
