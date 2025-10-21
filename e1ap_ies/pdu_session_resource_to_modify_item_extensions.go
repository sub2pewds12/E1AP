package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceToModifyItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToModifyItemExtensions struct {
	SNSSAI                                *SNSSAI
	CommonNetworkInstance                 *CommonNetworkInstance
	RedundantNGULUPTNLInformation         *UPTNLInformation
	RedundantCommonNetworkInstance        *CommonNetworkInstance
	DataForwardingtoEUTRANInformationList *DataForwardingtoEUTRANInformationList
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToModifyItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceToModifyItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToModifyItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceToModifyItemExtensions")
}
