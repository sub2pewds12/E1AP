package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// AdditionalPDCPduplicationInformation is a generated ENUMERATED type.
type AdditionalPDCPduplicationInformation struct {
	Value aper.Enumerated
}

const (
	AdditionalPDCPduplicationInformationThree aper.Enumerated = 0
	AdditionalPDCPduplicationInformationFour  aper.Enumerated = 1
)

func (e *AdditionalPDCPduplicationInformation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum AdditionalPDCPduplicationInformation to be generated here.
	return fmt.Errorf("Encode not implemented for enum AdditionalPDCPduplicationInformation")
}

func (e *AdditionalPDCPduplicationInformation) Decode(r *aper.AperReader) error {
	// Decode logic for enum AdditionalPDCPduplicationInformation to be generated here.
	return fmt.Errorf("Decode not implemented for enum AdditionalPDCPduplicationInformation")
}
