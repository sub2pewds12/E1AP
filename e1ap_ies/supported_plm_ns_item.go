package e1ap_ies

import (
	"fmt"
	"io"
)

// SupportedPLMNsItem is a generated SEQUENCE type.
type SupportedPLMNsItem struct {
	PLMNIdentity             PLMNIdentity                `aper:"lb:3,ub:3,mandatory,ignore,ext"`
	SliceSupportList         []SliceSupportItem          `aper:"optional,ext"`
	NRCGISupportList         []NRCGISupportItem          `aper:"optional,ext"`
	QOSParametersSupportList *QOSParametersSupportList   `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
	// Possible extensions:
	// - NPNSupportInfo (ID: id-NPNSupportInfo)
	// - ExtendedSliceSupportList (ID: id-ExtendedSliceSupportList)
	// - ExtendedNRCGISupportList (ID: id-Extended-NR-CGI-Support-List)
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItem) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for SupportedPLMNsItem")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SupportedPLMNsItem) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for SupportedPLMNsItem")
}
