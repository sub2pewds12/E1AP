package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceFailedItem is a generated SEQUENCE type.
type PDUSessionResourceFailedItem struct {
	PDUSessionID aper.Integer                `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
