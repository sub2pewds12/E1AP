package e1ap_ies

import (
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
