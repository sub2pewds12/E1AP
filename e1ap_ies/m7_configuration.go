package e1ap_ies

import "github.com/lvdund/ngap/aper"

// M7Configuration is a generated SEQUENCE type.
type M7Configuration struct {
	M7period     aper.Integer `aper:"lb:1,ub:60,mandatory,ext"`
	M7LinksToLog LinksToLog   `aper:"mandatory,ext"`
}
