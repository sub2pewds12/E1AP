package e1ap_ies

// SDAPConfiguration is a generated SEQUENCE type.
type SDAPConfiguration struct {
	DefaultDRB   DefaultDRB                  `aper:"mandatory,ext"`
	SDAPHeaderUL SDAPHeaderUL                `aper:"mandatory,ext"`
	SDAPHeaderDL SDAPHeaderDL                `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
