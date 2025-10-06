package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResetAcknowledge is a generated SEQUENCE type.
type ResetAcknowledge struct {
	TransactionID                             aper.Integer                `aper:"lb:0,ub:255,mandatory,reject,ext"`
	UEAssociatedLogicalE1ConnectionListResAck []ProtocolIESingleContainer `aper:"optional,ignore,ext"`
	CriticalityDiagnostics                    *CriticalityDiagnostics     `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResetAcknowledge) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ResetAcknowledge")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResetAcknowledge) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ResetAcknowledge")
}
