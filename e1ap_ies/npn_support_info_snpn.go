package e1ap_ies

// NPNSupportInfoSNPN is a generated SEQUENCE type.
type NPNSupportInfoSNPN struct {
	NID          NID                         `aper:"lb:44,ub:44,mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
