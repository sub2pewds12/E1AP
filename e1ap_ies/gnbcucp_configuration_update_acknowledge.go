package e1ap_ies

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPConfigurationUpdateAcknowledge is a generated SEQUENCE type.
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                aper.Integer                   `aper:"lb:0,ub:255,mandatory,reject,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics        `aper:"optional,ignore,ext"`
	GNBCUCPTNLASetupList         []GNBCUCPTNLASetupItem         `aper:"optional,ignore,ext"`
	GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo     `aper:"optional,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPConfigurationUpdateAcknowledge) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUCPConfigurationUpdateAcknowledge")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPConfigurationUpdateAcknowledge) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUCPConfigurationUpdateAcknowledge")
}
