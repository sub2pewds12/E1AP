package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBsSubjectToCounterCheckItemEUTRAN is a generated SEQUENCE type.
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID       aper.Integer `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount PDCPCount    `aper:"mandatory,ext"`
	PDCPDLCount PDCPCount    `aper:"mandatory,ext"`
}
