package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdateFailure is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdateFailure struct {
	TransactionID          aper.Integer            `aper:"lb:0,ub:255,mandatory,reject,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPConfigurationUpdateFailure) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUUPConfigurationUpdateFailure")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPConfigurationUpdateFailure) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUUPConfigurationUpdateFailure")
}
