package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GTPTLAItem is a generated SEQUENCE type.
type GTPTLAItem struct {
	GTPTransportLayerAddresses aper.BitString              `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	IEExtensions               *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GTPTLAItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GTPTLAItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GTPTLAItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GTPTLAItem")
}
