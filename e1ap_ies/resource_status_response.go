package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusResponse is a generated SEQUENCE type.
type ResourceStatusResponse struct {
	TransactionID          aper.Integer            `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPMeasurementID   aper.Integer            `aper:"lb:1,ub:4095,mandatory,reject,ext"`
	GNBCUUPMeasurementID   aper.Integer            `aper:"lb:1,ub:4095,mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusResponse) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ResourceStatusResponse")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusResponse) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ResourceStatusResponse")
}
