package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBStatusItem is a generated SEQUENCE type.
type DRBStatusItem struct {
	DRBID       aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPDLCount *PDCPCount   `aper:"optional,ext"`
	PDCPULCount *PDCPCount   `aper:"optional,ext"`
}
