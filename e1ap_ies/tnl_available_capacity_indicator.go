package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TNLAvailableCapacityIndicator is a generated SEQUENCE type.
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   aper.Integer `aper:"lb:0,ub:16777216,mandatory,ext"`
	DLTNLAvailableCapacity aper.Integer `aper:"lb:0,ub:100,mandatory,ext"`
	ULTNLOfferedCapacity   aper.Integer `aper:"lb:0,ub:16777216,mandatory,ext"`
	ULTNLAvailableCapacity aper.Integer `aper:"lb:0,ub:100,mandatory,ext"`
}
