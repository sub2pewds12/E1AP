package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsagePerPDUSessionReportSecondaryRATType From: 9_4_5_Information_Element_Definitions.txt:346
const (
	DataUsagePerPDUSessionReportSecondaryRATTypeNR    aper.Enumerated = 0
	DataUsagePerPDUSessionReportSecondaryRATTypeEUTRA aper.Enumerated = 1
)

type DataUsagePerPDUSessionReportSecondaryRATType struct {
	Value aper.Enumerated
}
