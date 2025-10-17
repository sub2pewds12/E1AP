package e1ap_ies

// UplinkOnlyROHC is a generated SEQUENCE type.
type UplinkOnlyROHC struct {
	MaxCID       UplinkOnlyROHCMaxCID        `aper:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles UplinkOnlyROHCROHCProfiles  `aper:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC *UplinkOnlyROHCContinueROHC `aper:"optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
