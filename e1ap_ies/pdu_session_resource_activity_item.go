package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceActivityItem From: 9_4_5_Information_Element_Definitions.txt:1655
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceActivityItem struct {
	PDUSessionID               aper.Integer               `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionResourceActivity PDUSessionResourceActivity `aper:"mandatory,ext"`
}
