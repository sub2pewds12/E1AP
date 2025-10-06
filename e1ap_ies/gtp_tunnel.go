package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GTPTunnel is a generated SEQUENCE type.
type GTPTunnel struct {
	TransportLayerAddress aper.BitString   `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	GTPTEID               aper.OctetString `aper:"lb:4,ub:4,mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GTPTunnel) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GTPTunnel")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GTPTunnel) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GTPTunnel")
}
