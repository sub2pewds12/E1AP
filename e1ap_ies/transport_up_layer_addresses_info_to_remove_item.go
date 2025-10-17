package e1ap_ies

import (
	"fmt"
	"io"
)

// TransportUPLayerAddressesInfoToRemoveItem is a generated SEQUENCE type.
type TransportUPLayerAddressesInfoToRemoveItem struct {
	IPSecTransportLayerAddress         TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	GTPTransportLayerAddressesToRemove []GTPTLAItem                `aper:"ub:MaxnoofGTPTLAs,optional,ext"`
	IEExtensions                       *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TransportUPLayerAddressesInfoToRemoveItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for TransportUPLayerAddressesInfoToRemoveItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TransportUPLayerAddressesInfoToRemoveItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for TransportUPLayerAddressesInfoToRemoveItem")
}
