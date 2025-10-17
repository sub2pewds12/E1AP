package e1ap_ies

// PDUSessionResourceFailedItem is a generated SEQUENCE type.
type PDUSessionResourceFailedItem struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
