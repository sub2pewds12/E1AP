package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingInformationExtensions is a generated type-safe wrapper for extensions.
type DataForwardingInformationExtensions struct {
	DataForwardingtoNGRANQoSFlowInformationList *DataForwardingtoNGRANQoSFlowInformationList
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformationExtensions) Encode(w *aper.AperWriter) error {
	// TODO: Implement the complex APER extension container encoding logic.
	// This involves creating a bitmap of present extensions and encoding each one as an open type.
	return fmt.Errorf("Encode not yet implemented for DataForwardingInformationExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingInformationExtensions) Decode(r *aper.AperReader) error {
	// TODO: Implement the complex APER extension container decoding logic.
	// This involves reading a bitmap and then decoding each present extension as an open type.
	return fmt.Errorf("Decode not yet implemented for DataForwardingInformationExtensions")
}
