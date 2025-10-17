package e1ap_ies

// PacketErrorRate is a generated SEQUENCE type.
type PacketErrorRate struct {
	PERScalar    PERScalar                   `aper:"lb:0,ub:9,mandatory,ext"`
	PERExponent  PERExponent                 `aper:"lb:0,ub:9,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
