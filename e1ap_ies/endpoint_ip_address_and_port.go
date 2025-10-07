package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// EndpointIPAddressAndPort is a generated SEQUENCE type.
type EndpointIPAddressAndPort struct {
	EndpointIPAddress aper.BitString              `aper:"lb:1,ub:160,mandatory,ignore"`
	PortNumber        aper.BitString              `aper:"lb:16,ub:16,mandatory"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EndpointIPAddressAndPort) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for EndpointIPAddressAndPort")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EndpointIPAddressAndPort) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for EndpointIPAddressAndPort")
}
