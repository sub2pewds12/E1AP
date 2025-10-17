package e1ap_ies

// GTPTLAItem is a generated SEQUENCE type.
type GTPTLAItem struct {
	GTPTransportLayerAddresses TransportLayerAddress       `aper:"lb:1,ub:160,mandatory"`
	IEExtensions               *ProtocolExtensionContainer `aper:"optional"`
}
