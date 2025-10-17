package e1ap_ies

// NPNContextInfoSNPN is a generated SEQUENCE type.
type NPNContextInfoSNPN struct {
	NID          NID                         `aper:"lb:44,ub:44,mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
