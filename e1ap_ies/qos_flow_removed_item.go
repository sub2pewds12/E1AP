package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QOSFlowRemovedItem From: 9_4_5_Information_Element_Definitions.txt:2037
// ASN.1 Data Type: SEQUENCE
type QOSFlowRemovedItem struct {
	QOSFlowIdentifier             aper.Integer                                `aper:"lb:0,ub:63,mandatory,ext"`
	QOSFlowReleasedInSession      *QOSFlowRemovedItemQOSFlowReleasedInSession `aper:"optional,ext"`
	QOSFlowAccumulatedSessionTime *aper.OctetString                           `aper:"lb:5,ub:5,optional,ext"`
}
