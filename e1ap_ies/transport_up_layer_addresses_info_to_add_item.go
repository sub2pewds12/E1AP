package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// TransportUPLayerAddressesInfoToAddItem is a generated SEQUENCE type.
type TransportUPLayerAddressesInfoToAddItem struct {
	IPSecTransportLayerAddress      aper.BitString `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	GTPTransportLayerAddressesToAdd []GTPTLAItem   `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportUPLayerAddressesInfoToAddItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for TransportUPLayerAddressesInfoToAddItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransportUPLayerAddressesInfoToAddItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for TransportUPLayerAddressesInfoToAddItem")
}
