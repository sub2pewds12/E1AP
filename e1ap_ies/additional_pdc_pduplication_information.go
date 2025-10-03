package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// AdditionalPDCPduplicationInformation From: 9_4_5_Information_Element_Definitions.txt:103
// ASN.1 Data Type: ENUMERATED
const (
	AdditionalPDCPduplicationInformationThree aper.Enumerated = 0
	AdditionalPDCPduplicationInformationFour  aper.Enumerated = 1
)

type AdditionalPDCPduplicationInformation struct {
	Value aper.Enumerated
}
