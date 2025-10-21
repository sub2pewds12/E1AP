package e1ap_ies

import (
	"fmt"
	"io"
)

// DataForwardingInformationExtensions is a generated type-safe wrapper for extensions.
type DataForwardingInformationExtensions struct {
	DataForwardingtoNGRANQoSFlowInformationList *DataForwardingtoNGRANQoSFlowInformationList
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformationExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for DataForwardingInformationExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingInformationExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for DataForwardingInformationExtensions")
}
