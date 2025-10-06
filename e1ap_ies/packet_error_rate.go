package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PacketErrorRate is a generated SEQUENCE type.
type PacketErrorRate struct {
	PERScalar   aper.Integer `aper:"lb:0,ub:9,mandatory,ext"`
	PERExponent aper.Integer `aper:"lb:0,ub:9,mandatory,ext"`
}
