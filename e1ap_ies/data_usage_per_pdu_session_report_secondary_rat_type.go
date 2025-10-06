package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataUsagePerPDUSessionReportSecondaryRATType is a generated ENUMERATED type.
type DataUsagePerPDUSessionReportSecondaryRATType struct {
	Value aper.Enumerated
}

const (
	DataUsagePerPDUSessionReportSecondaryRATTypeNR    aper.Enumerated = 0
	DataUsagePerPDUSessionReportSecondaryRATTypeEUTRA aper.Enumerated = 1
)
