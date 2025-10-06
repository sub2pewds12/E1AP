package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataUsagePerQOSFlowItemSecondaryRATType is a generated ENUMERATED type.
type DataUsagePerQOSFlowItemSecondaryRATType struct {
	Value aper.Enumerated
}

const (
	DataUsagePerQOSFlowItemSecondaryRATTypeNR    aper.Enumerated = 0
	DataUsagePerQOSFlowItemSecondaryRATTypeEUTRA aper.Enumerated = 1
)
