package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// ConfidentialityProtectionResult is a generated ENUMERATED type.
type ConfidentialityProtectionResult struct {
	Value int64
}

const (
	ConfidentialityProtectionResultPerformed    int64 = 0
	ConfidentialityProtectionResultNotPerformed int64 = 1
)

// Encode implements the MessageEncoder interface for ConfidentialityProtectionResult.
func (e *ConfidentialityProtectionResult) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for ConfidentialityProtectionResult.
func (e *ConfidentialityProtectionResult) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
