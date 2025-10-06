package e1ap_ies

import "github.com/lvdund/ngap/aper"

// HWCapacityIndicator is a generated SEQUENCE type.
type HWCapacityIndicator struct {
	OfferedThroughput   aper.Integer `aper:"lb:1,ub:16777216,mandatory,ext"`
	AvailableThroughput aper.Integer `aper:"lb:0,ub:100,mandatory,ext"`
}
