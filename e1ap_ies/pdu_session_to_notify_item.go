package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionToNotifyItem From: 9_4_5_Information_Element_Definitions.txt:1888
// ASN.1 Data Type: SEQUENCE
type PDUSessionToNotifyItem struct {
	PDUSessionID aper.Integer  `aper:"lb:0,ub:255,mandatory,ext"`
	QOSFlowList  []QOSFlowItem `aper:"mandatory,ext"`
}
