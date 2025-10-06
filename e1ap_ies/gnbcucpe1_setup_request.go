package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPE1SetupRequest is a generated SEQUENCE type.
type GNBCUCPE1SetupRequest struct {
	TransactionID             aper.Integer               `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPE1SetupRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUCPE1SetupRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPE1SetupRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUCPE1SetupRequest")
}
