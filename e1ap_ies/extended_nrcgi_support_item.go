package e1ap_ies

// ExtendedNRCGISupportItem is a generated SEQUENCE type.
type ExtendedNRCGISupportItem struct {
	NRCGI        NRCGI                       `aper:"mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
