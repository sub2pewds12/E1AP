package e1ap_ies

// PDUSessionResourceFailedToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceFailedToModifyItem struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
