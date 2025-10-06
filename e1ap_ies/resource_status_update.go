package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusUpdate is a generated SEQUENCE type.
type ResourceStatusUpdate struct {
	TransactionID                 aper.Integer                   `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPMeasurementID          aper.Integer                   `aper:"lb:1,ub:4095,mandatory,reject,ext"`
	GNBCUUPMeasurementID          *aper.Integer                  `aper:"lb:1,ub:4095,optional,ignore,ext"`
	TNLAvailableCapacityIndicator *TNLAvailableCapacityIndicator `aper:"optional,ignore,ext"`
	HWCapacityIndicator           HWCapacityIndicator            `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusUpdate) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ResourceStatusUpdate")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusUpdate) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ResourceStatusUpdate")
}
