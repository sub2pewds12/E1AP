package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionToNotifyItem is a generated SEQUENCE type.
type PDUSessionToNotifyItem struct {
	PDUSessionID aper.Integer  `aper:"lb:0,ub:255,mandatory,ext"`
	QOSFlowList  []QOSFlowItem `aper:"mandatory,ext"`
}
