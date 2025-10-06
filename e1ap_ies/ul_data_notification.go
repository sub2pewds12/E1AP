package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ULDataNotification is a generated SEQUENCE type.
type ULDataNotification struct {
	GNBCUCPUEE1APID        aper.Integer             `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID        aper.Integer             `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	PDUSessionToNotifyList []PDUSessionToNotifyItem `aper:"mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ULDataNotification) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ULDataNotification")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ULDataNotification) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ULDataNotification")
}
