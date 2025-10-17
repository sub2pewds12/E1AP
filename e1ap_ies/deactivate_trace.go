package e1ap_ies

import (
	"fmt"
	"io"
)

// DeactivateTrace is a generated SEQUENCE type.
type DeactivateTrace struct {
	GNBCUCPUEE1APID GNBCUCPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID GNBCUUPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceID         TraceID         `aper:"lb:8,ub:8,mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DeactivateTrace) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for DeactivateTrace")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DeactivateTrace) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for DeactivateTrace")
}
