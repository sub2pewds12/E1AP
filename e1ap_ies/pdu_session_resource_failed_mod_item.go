package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceFailedModItem From: 9_4_5_Information_Element_Definitions.txt:1695
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceFailedModItem struct {
	PDUSessionID aper.Integer `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause        `aper:"mandatory,ignore,ext"`
}
