package e1ap_ies

// TReorderingTimer is a generated SEQUENCE type.
type TReorderingTimer struct {
	TReordering  TReordering                 `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
