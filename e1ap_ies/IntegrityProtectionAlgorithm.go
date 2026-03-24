package e1ap_ies

import (
	"asn1go/per"
)

// IntegrityProtectionAlgorithm is a generated ENUMERATED type.
type IntegrityProtectionAlgorithm struct {
	Value int64
}

const (
	IntegrityProtectionAlgorithmNIA0     int64 = 0
	IntegrityProtectionAlgorithmI128NIA1 int64 = 1
	IntegrityProtectionAlgorithmI128NIA2 int64 = 2
	IntegrityProtectionAlgorithmI128NIA3 int64 = 3
)

// Encode implements the MessageEncoder interface for IntegrityProtectionAlgorithm.
func (e *IntegrityProtectionAlgorithm) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for IntegrityProtectionAlgorithm.
func (e *IntegrityProtectionAlgorithm) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
