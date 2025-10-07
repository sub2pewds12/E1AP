package e1ap_ies

// DAPSRequestInfo is a generated SEQUENCE type.
type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator `aper:"mandatory,ext"`
	IEExtensions  *ProtocolExtensionContainer  `aper:"optional,ext"`
}
