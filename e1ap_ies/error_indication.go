package e1ap_ies

import (
	"fmt"
	"io"
)

// ErrorIndication is a generated SEQUENCE type.
type ErrorIndication struct {
	TransactionID          TransactionID           `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPUEE1APID        *GNBCUCPUEE1APID        `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	GNBCUUPUEE1APID        *GNBCUUPUEE1APID        `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	Cause                  *Cause                  `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ErrorIndication) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ErrorIndication")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ErrorIndication) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ErrorIndication")
}
