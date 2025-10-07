package e1ap_ies

// DLDiscarding is a generated SEQUENCE type.
type DLDiscarding struct {
	DLDiscardingCountVal PDCPCount                   `aper:"mandatory"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional"`
}
