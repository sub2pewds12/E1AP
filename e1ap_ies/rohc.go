package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ROHC is a generated SEQUENCE type.
type ROHC struct {
	MaxCID       aper.Integer      `aper:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles aper.Integer      `aper:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC *ROHCContinueROHC `aper:"optional,ext"`
}
