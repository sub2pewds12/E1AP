package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceFailedToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceFailedToModifyItem struct {
	PDUSessionID aper.Integer                `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
