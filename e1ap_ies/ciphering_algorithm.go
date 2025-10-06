package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CipheringAlgorithm is a generated ENUMERATED type.
type CipheringAlgorithm struct {
	Value aper.Enumerated
}

const (
	CipheringAlgorithmNEA0     aper.Enumerated = 0
	CipheringAlgorithmC128NEA1 aper.Enumerated = 1
	CipheringAlgorithmC128NEA2 aper.Enumerated = 2
	CipheringAlgorithmC128NEA3 aper.Enumerated = 3
)
