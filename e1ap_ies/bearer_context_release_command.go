package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextReleaseCommand is a generated SEQUENCE type.
type BearerContextReleaseCommand struct {
	GNBCUCPUEE1APID aper.Integer `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	Cause           Cause        `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextReleaseCommand) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextReleaseCommand")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextReleaseCommand) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextReleaseCommand")
}
