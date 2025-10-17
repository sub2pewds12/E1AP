package e1ap_ies

import (
	"fmt"
	"io"
)

// GNBCUUPCounterCheckRequest is a generated SEQUENCE type.
type GNBCUUPCounterCheckRequest struct {
	GNBCUCPUEE1APID                  GNBCUCPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID                  GNBCUUPUEE1APID                  `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	SystemGNBCUUPCounterCheckRequest SystemGNBCUUPCounterCheckRequest `aper:"mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPCounterCheckRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUUPCounterCheckRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPCounterCheckRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUUPCounterCheckRequest")
}
