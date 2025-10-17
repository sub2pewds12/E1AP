package e1ap_ies

// ROHC is a generated SEQUENCE type.
type ROHC struct {
	MaxCID       ROHCMaxCID                  `aper:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles ROHCROHCProfiles            `aper:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC *ROHCContinueROHC           `aper:"optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
