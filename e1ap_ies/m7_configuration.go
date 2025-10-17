package e1ap_ies

// M7Configuration is a generated SEQUENCE type.
type M7Configuration struct {
	M7period     M7period                    `aper:"lb:1,ub:60,mandatory,ext"`
	M7LinksToLog LinksToLog                  `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
