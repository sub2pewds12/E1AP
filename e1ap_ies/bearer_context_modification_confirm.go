package e1ap_ies

import (
	"fmt"
	"io"
)

// BearerContextModificationConfirm is a generated SEQUENCE type.
type BearerContextModificationConfirm struct {
	GNBCUCPUEE1APID                        GNBCUCPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                        GNBCUUPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationConfirm *SystemBearerContextModificationConfirm `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextModificationConfirm) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextModificationConfirm")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextModificationConfirm) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextModificationConfirm")
}
