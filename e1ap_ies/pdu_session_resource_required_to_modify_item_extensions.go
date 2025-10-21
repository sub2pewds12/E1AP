package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceRequiredToModifyItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceRequiredToModifyItemExtensions struct {
	RedundantNGDLUPTNLInformation *UPTNLInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceRequiredToModifyItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceRequiredToModifyItemExtensions")
}
