package e1ap_ies

import (
	"fmt"
	"io"
)

// ResourceStatusResponse is a generated SEQUENCE type.
type ResourceStatusResponse struct {
	TransactionID          TransactionID                                   `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPMeasurementID   ResourceStatusResponseIEsIDGNBCUCPMeasurementID `aper:"lb:1,ub:4095,mandatory,reject,ext"`
	GNBCUUPMeasurementID   ResourceStatusResponseIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics                         `aper:"optional,ignore,ext"`
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
