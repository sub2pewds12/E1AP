package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// DLDataNotification is a generated SEQUENCE type.
type DLDataNotification struct {
	GNBCUCPUEE1APID aper.Integer  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID aper.Integer  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PPI             *aper.Integer `aper:"lb:0,ub:7,optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DLDataNotification) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for DLDataNotification")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DLDataNotification) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for DLDataNotification")
}
