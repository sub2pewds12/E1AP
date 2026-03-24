package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// ConfidentialityProtectionIndication is a generated ENUMERATED type.
type ConfidentialityProtectionIndication struct {
	Value int64
}

const (
	ConfidentialityProtectionIndicationRequired  int64 = 0
	ConfidentialityProtectionIndicationPreferred int64 = 1
	ConfidentialityProtectionIndicationNotNeeded int64 = 2
)

// Encode implements the MessageEncoder interface for ConfidentialityProtectionIndication.
func (e *ConfidentialityProtectionIndication) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for ConfidentialityProtectionIndication.
func (e *ConfidentialityProtectionIndication) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
