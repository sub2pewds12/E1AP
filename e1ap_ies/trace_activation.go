package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// TraceActivation is a generated SEQUENCE type.
type TraceActivation struct {
	TraceID                        aper.OctetString `aper:"lb:8,ub:8,mandatory,ignore,ext"`
	InterfacesToTrace              aper.BitString   `aper:"lb:8,ub:8,mandatory,ext"`
	TraceDepth                     TraceDepth       `aper:"mandatory,ext"`
	TraceCollectionEntityIPAddress aper.BitString   `aper:"lb:1,ub:160,mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceActivation) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for TraceActivation")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceActivation) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for TraceActivation")
}
