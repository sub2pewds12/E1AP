package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusRequest is a generated SEQUENCE type.
type ResourceStatusRequest struct {
	TransactionID         aper.Integer          `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPMeasurementID  aper.Integer          `aper:"lb:1,ub:4095,mandatory,reject,ext"`
	GNBCUUPMeasurementID  *aper.Integer         `aper:"lb:1,ub:4095,optional,ignore,ext"`
	RegistrationRequest   RegistrationRequest   `aper:"mandatory,reject,ext"`
	ReportCharacteristics *aper.BitString       `aper:"lb:36,ub:36,conditional,reject,ext"`
	ReportingPeriodicity  *ReportingPeriodicity `aper:"optional,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for ResourceStatusRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for ResourceStatusRequest")
}
