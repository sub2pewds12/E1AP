package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UPParametersItem is a generated SEQUENCE type.
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation `aper:"mandatory,ext"`
	CellGroupID      aper.Integer     `aper:"lb:0,ub:3,mandatory,ext"`
}
