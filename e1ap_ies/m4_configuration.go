package e1ap_ies

// M4Configuration is a generated SEQUENCE type.
type M4Configuration struct {
	M4period     M4period                    `aper:"mandatory,ext"`
	M4LinksToLog LinksToLog                  `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
