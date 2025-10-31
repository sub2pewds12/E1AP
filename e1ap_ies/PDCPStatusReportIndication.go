package e1ap_ies

import (
	"fmt"

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

func (e *PDCPStatusReportIndication) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDCPStatusReportIndication to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDCPStatusReportIndication")
}

func (e *PDCPStatusReportIndication) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDCPStatusReportIndication to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDCPStatusReportIndication")
}
