package e1ap_ies

// EHCCommonParameters is a generated SEQUENCE type.
type EHCCommonParameters struct {
	EhcCIDLength EHCCommonParametersEhcCIDLength `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer     `aper:"optional,ext"`
}
