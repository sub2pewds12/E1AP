package e1ap_ies

import (
	"fmt"
	"io"
)

// ULUPTNLAddressToUpdateItem is a generated SEQUENCE type.
type ULUPTNLAddressToUpdateItem struct {
	OldTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	NewTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ignore,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ULUPTNLAddressToUpdateItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ULUPTNLAddressToUpdateItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ULUPTNLAddressToUpdateItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ULUPTNLAddressToUpdateItem")
}
