package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// CipheringAlgorithm is a generated ENUMERATED type.
type CipheringAlgorithm struct {
	Value int64
}

const (
	CipheringAlgorithmNEA0     int64 = 0
	CipheringAlgorithmC128NEA1 int64 = 1
	CipheringAlgorithmC128NEA2 int64 = 2
	CipheringAlgorithmC128NEA3 int64 = 3
)

// Encode implements the MessageEncoder interface for CipheringAlgorithm.
func (e *CipheringAlgorithm) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CipheringAlgorithm.
func (e *CipheringAlgorithm) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
