package e1ap_ies

// SliceSupportItem is a generated SEQUENCE type.
type SliceSupportItem struct {
	SNSSAI       SNSSAI                      `aper:"mandatory"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional"`
}
