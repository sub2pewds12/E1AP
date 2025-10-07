package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceConfirmModifiedItem is a generated SEQUENCE type.
type PDUSessionResourceConfirmModifiedItem struct {
	PDUSessionID                aper.Integer                  `aper:"lb:0,ub:255,mandatory,ext"`
	DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN `aper:"optional,ext"`
	IEExtensions                *ProtocolExtensionContainer   `aper:"optional,ext"`
}
