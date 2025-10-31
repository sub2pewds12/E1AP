package e1ap_ies

import (
	"fmt"

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

func (e *CipheringAlgorithm) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CipheringAlgorithm to be generated here.
	return fmt.Errorf("Encode not implemented for enum CipheringAlgorithm")
}

func (e *CipheringAlgorithm) Decode(r *aper.AperReader) error {
	// Decode logic for enum CipheringAlgorithm to be generated here.
	return fmt.Errorf("Decode not implemented for enum CipheringAlgorithm")
}
