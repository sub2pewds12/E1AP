package e1ap_ies

import (
	"fmt"
	"io"
)

// DLUPTNLAddressToUpdateItem is a generated SEQUENCE type.
type DLUPTNLAddressToUpdateItem struct {
	OldTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	NewTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DLUPTNLAddressToUpdateItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for DLUPTNLAddressToUpdateItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DLUPTNLAddressToUpdateItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for DLUPTNLAddressToUpdateItem")
}
