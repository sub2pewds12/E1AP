package e1ap_ies

import (
	"fmt"
	"io"
)

// DLDataNotification is a generated SEQUENCE type.
type DLDataNotification struct {
	GNBCUCPUEE1APID        GNBCUCPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        GNBCUUPUEE1APID          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PPI                    *PPI                     `aper:"lb:0,ub:7,optional,ignore,ext"`
	PDUSessionToNotifyList []PDUSessionToNotifyItem `aper:"optional,ignore,ext"`
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
