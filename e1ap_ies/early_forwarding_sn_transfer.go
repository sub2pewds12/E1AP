package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingSNTransfer is a generated SEQUENCE type.
type EarlyForwardingSNTransfer struct {
	GNBCUCPUEE1APID                  aper.Integer                       `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  aper.Integer                       `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DRBsSubjectToEarlyForwardingList []DRBsSubjectToEarlyForwardingItem `aper:"mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EarlyForwardingSNTransfer) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for EarlyForwardingSNTransfer")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EarlyForwardingSNTransfer) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for EarlyForwardingSNTransfer")
}
