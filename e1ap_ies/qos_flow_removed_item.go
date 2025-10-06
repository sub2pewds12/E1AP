package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowRemovedItem is a generated SEQUENCE type.
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             aper.Integer                                `aper:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession `aper:"optional,ext"`
	QOSFlowAccumulatedSessionTime *aper.OctetString                           `aper:"lb:5,ub:5,optional,ext"`
}
