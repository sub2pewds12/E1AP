package e1ap_ies

import (
	"fmt"
	"io"
)

// BearerContextModificationResponse is a generated SEQUENCE type.
type BearerContextModificationResponse struct {
	GNBCUCPUEE1APID                         GNBCUCPUEE1APID                          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                         GNBCUUPUEE1APID                          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationResponse *SystemBearerContextModificationResponse `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextModificationResponse) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextModificationResponse")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextModificationResponse) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextModificationResponse")
}
