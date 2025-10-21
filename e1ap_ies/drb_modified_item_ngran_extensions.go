package e1ap_ies

import (
	"fmt"
	"io"
)

// DRBModifiedItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBModifiedItemNGRANExtensions struct {
	EarlyForwardingCOUNTInfo         *EarlyForwardingCOUNTInfo
	OldQoSFlowMapULendmarkerexpected *QOSFlowList
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemNGRANExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for DRBModifiedItemNGRANExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBModifiedItemNGRANExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for DRBModifiedItemNGRANExtensions")
}
