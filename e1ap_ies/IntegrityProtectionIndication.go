package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// IntegrityProtectionIndication is a generated ENUMERATED type.
type IntegrityProtectionIndication struct {
	Value int64
}

const (
	IntegrityProtectionIndicationRequired  int64 = 0
	IntegrityProtectionIndicationPreferred int64 = 1
	IntegrityProtectionIndicationNotNeeded int64 = 2
)

// Encode implements the MessageEncoder interface for IntegrityProtectionIndication.
func (e *IntegrityProtectionIndication) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for IntegrityProtectionIndication.
func (e *IntegrityProtectionIndication) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
