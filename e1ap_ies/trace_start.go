package e1ap_ies

import (
	"fmt"
	"io"
)

// TraceStart is a generated SEQUENCE type.
type TraceStart struct {
	GNBCUCPUEE1APID GNBCUCPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID GNBCUUPUEE1APID `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	TraceActivation TraceActivation `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceStart) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for TraceStart")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceStart) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for TraceStart")
}
