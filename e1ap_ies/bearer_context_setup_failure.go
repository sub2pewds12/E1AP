package e1ap_ies

import (
	"fmt"
	"io"
)

// BearerContextSetupFailure is a generated SEQUENCE type.
type BearerContextSetupFailure struct {
	GNBCUCPUEE1APID        GNBCUCPUEE1APID         `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        *GNBCUUPUEE1APID        `aper:"lb:0,ub:4294967295,optional,ignore,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextSetupFailure) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextSetupFailure")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextSetupFailure) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextSetupFailure")
}
