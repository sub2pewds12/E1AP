package e1ap_ies

import (
	"fmt"
	"io"
)

// GBRQoSFlowInformationExtensions is a generated type-safe wrapper for extensions.
type GBRQoSFlowInformationExtensions struct {
	AlternativeQoSParaSetList *AlternativeQoSParaSetList
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQoSFlowInformationExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for GBRQoSFlowInformationExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GBRQoSFlowInformationExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for GBRQoSFlowInformationExtensions")
}
