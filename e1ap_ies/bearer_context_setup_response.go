package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextSetupResponse is a generated SEQUENCE type.
type BearerContextSetupResponse struct {
	GNBCUCPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  aper.Integer                     `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemBearerContextSetupResponse SystemBearerContextSetupResponse `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextSetupResponse) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextSetupResponse")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextSetupResponse) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextSetupResponse")
}
