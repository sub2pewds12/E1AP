package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowItem is a generated SEQUENCE type.
type QOSFlowItem struct {
	QOSFlowIdentifier aper.Integer `aper:"lb:0,ub:63,mandatory,ext"`
}
