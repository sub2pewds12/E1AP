package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDCPStatusReportIndication is a generated ENUMERATED type.
type PDCPStatusReportIndication struct {
	Value aper.Enumerated
}

const (
	PDCPStatusReportIndicationDownlink aper.Enumerated = 0
	PDCPStatusReportIndicationUplink   aper.Enumerated = 1
	PDCPStatusReportIndicationBoth     aper.Enumerated = 2
)
