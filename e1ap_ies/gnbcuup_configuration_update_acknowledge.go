package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUUPConfigurationUpdateAcknowledge struct {
	TransactionID             aper.Integer               `aper:"lb:0,ub:255,mandatory,reject,ext"`
	CriticalityDiagnostics    *CriticalityDiagnostics    `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPConfigurationUpdateAcknowledge) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUUPConfigurationUpdateAcknowledge")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPConfigurationUpdateAcknowledge) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUUPConfigurationUpdateAcknowledge")
}
