package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// BearerContextReleaseRequest is a generated SEQUENCE type.
type BearerContextReleaseRequest struct {
	GNBCUCPUEE1APID aper.Integer    `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer    `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBStatusList   []DRBStatusItem `aper:"lb:1,ub:MaxnoofDRBs,optional,ignore,ext"`
	Cause           Cause           `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextReleaseRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for BearerContextReleaseRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextReleaseRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for BearerContextReleaseRequest")
}
