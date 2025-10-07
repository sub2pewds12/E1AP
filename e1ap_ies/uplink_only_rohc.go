package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UplinkOnlyROHC is a generated SEQUENCE type.
type UplinkOnlyROHC struct {
	MaxCID       aper.Integer                `aper:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles aper.Integer                `aper:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC *UplinkOnlyROHCContinueROHC `aper:"optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
