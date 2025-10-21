package e1ap_ies

import (
	"fmt"
	"io"
)

// DRBToModifyItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToModifyItemNGRANExtensions struct {
	OldQoSFlowMapULendmarkerexpected *QOSFlowList
	DRBQOS                           *QoSFlowLevelQoSParameters
	EarlyForwardingCOUNTReq          *EarlyForwardingCOUNTReq
	EarlyForwardingCOUNTInfo         *EarlyForwardingCOUNTInfo
	DAPSRequestInfo                  *DAPSRequestInfo
	EarlyDataForwardingIndicator     *EarlyDataForwardingIndicator
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToModifyItemNGRANExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for DRBToModifyItemNGRANExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToModifyItemNGRANExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for DRBToModifyItemNGRANExtensions")
}
