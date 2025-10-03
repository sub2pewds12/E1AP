package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceFailedToModifyItem From: 9_4_5_Information_Element_Definitions.txt:1708
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceFailedToModifyItem struct {
	PDUSessionID aper.Integer `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause        `aper:"mandatory,ignore,ext"`
}
