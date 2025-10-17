package e1ap_ies

import (
	"fmt"
	"io"
)

// BearerContextModificationRequired is a generated SEQUENCE type.
type BearerContextModificationRequired struct {
	GNBCUCPUEE1APID                         GNBCUCPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                         GNBCUUPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextModificationRequired SystemBearerContextModificationRequired `aper:"mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextModificationRequired) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextModificationRequired")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextModificationRequired) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextModificationRequired")
}
