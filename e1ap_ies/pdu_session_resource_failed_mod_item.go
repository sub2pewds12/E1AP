package e1ap_ies

// PDUSessionResourceFailedModItem is a generated SEQUENCE type.
type PDUSessionResourceFailedModItem struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
