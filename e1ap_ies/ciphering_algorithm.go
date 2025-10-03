package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CipheringAlgorithm From: 9_4_5_Information_Element_Definitions.txt:234
// ASN.1 Data Type: ENUMERATED
const (
	CipheringAlgorithmNEA0     aper.Enumerated = 0
	CipheringAlgorithmC128NEA1 aper.Enumerated = 1
	CipheringAlgorithmC128NEA2 aper.Enumerated = 2
	CipheringAlgorithmC128NEA3 aper.Enumerated = 3
)

type CipheringAlgorithm struct {
	Value aper.Enumerated
}
