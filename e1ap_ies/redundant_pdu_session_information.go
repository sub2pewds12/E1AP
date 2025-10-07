package e1ap_ies

// RedundantPDUSessionInformation is a generated SEQUENCE type.
type RedundantPDUSessionInformation struct {
	RSN          RSN                         `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
