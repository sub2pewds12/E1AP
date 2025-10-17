package e1ap_ies

// UPSecuritykey is a generated SEQUENCE type.
type UPSecuritykey struct {
	EncryptionKey          EncryptionKey               `aper:"mandatory,ext"`
	IntegrityProtectionKey *IntegrityProtectionKey     `aper:"optional,ext"`
	IEExtensions           *ProtocolExtensionContainer `aper:"optional,ext"`
}
