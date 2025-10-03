package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceConfirmModifiedItem From: 9_4_5_Information_Element_Definitions.txt:1669
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceConfirmModifiedItem struct {
	PDUSessionID                aper.Integer                  `aper:"lb:0,ub:255,mandatory,ext"`
	DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN `aper:"optional,ext"`
}
