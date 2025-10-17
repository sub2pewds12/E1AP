package e1ap_ies

import (
	"fmt"
	"io"
)

// NRCGI is a generated SEQUENCE type.
type NRCGI struct {
	PLMNIdentity   PLMNIdentity                `aper:"lb:3,ub:3,mandatory,ignore"`
	NRCellIdentity NRCellIdentity              `aper:"lb:36,ub:36,mandatory"`
	IEExtensions   *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *NRCGI) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for NRCGI")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NRCGI) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for NRCGI")
}
