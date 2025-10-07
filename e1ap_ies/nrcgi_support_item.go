package e1ap_ies

// NRCGISupportItem is a generated SEQUENCE type.
type NRCGISupportItem struct {
	NRCGI        NRCGI                       `aper:"mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
