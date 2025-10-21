package e1ap_ies

import (
	"fmt"
	"io"
)

// PDCPConfigurationExtensions is a generated type-safe wrapper for extensions.
type PDCPConfigurationExtensions struct {
	PDCPStatusReportIndication           *PDCPStatusReportIndication
	AdditionalPDCPduplicationInformation *AdditionalPDCPduplicationInformation
	EHCParameters                        *EHCParameters
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPConfigurationExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDCPConfigurationExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDCPConfigurationExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDCPConfigurationExtensions")
}
