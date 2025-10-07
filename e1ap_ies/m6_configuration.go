package e1ap_ies

// M6Configuration is a generated SEQUENCE type.
type M6Configuration struct {
	M6reportInterval M6reportInterval            `aper:"mandatory,ext"`
	M6LinksToLog     LinksToLog                  `aper:"mandatory,ext"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional,ext"`
}
